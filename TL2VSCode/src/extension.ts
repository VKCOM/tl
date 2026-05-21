import * as vscode from 'vscode';

const PRIMITIVES = new Set([
    'int32', 'uint32', 'int64', 'uint64', 'float32', 'float64',
    'byte', 'string', 'bool', 'bit', 'true',
]);

export function activate(context: vscode.ExtensionContext): void {
    const indexer = new Tl2Indexer();
    const diagColl = vscode.languages.createDiagnosticCollection('tl2');
    context.subscriptions.push(diagColl);

    void indexer.ensureWorkspaceIndexed().then(() => refreshAllDiagnostics(indexer, diagColl));

    context.subscriptions.push(
        vscode.workspace.onDidOpenTextDocument((doc) => {
            if (doc.languageId === 'tl2') {
                void lintDocument(doc, indexer, diagColl);
            }
        }),
        vscode.workspace.onDidCloseTextDocument((doc) => {
            if (doc.languageId === 'tl2') {
                diagColl.delete(doc.uri);
            }
        }),
        vscode.workspace.onDidChangeTextDocument((e) => {
            if (e.document.languageId !== 'tl2') {
                return;
            }
            indexer.invalidate(e.document.uri);
            void lintDocument(e.document, indexer, diagColl);
        }),
        vscode.workspace.onDidChangeConfiguration((e) => {
            if (e.affectsConfiguration('tl2.diagnostics.unknownTypes')) {
                void refreshAllDiagnostics(indexer, diagColl);
            }
        }),
    );

    const watcher = vscode.workspace.createFileSystemWatcher('**/*.tl2');
    context.subscriptions.push(
        watcher,
        watcher.onDidChange((uri) => {
            indexer.invalidate(uri);
            void refreshAllDiagnostics(indexer, diagColl);
        }),
        watcher.onDidDelete((uri) => {
            indexer.invalidate(uri);
            indexer.refreshWorkspaceFiles();
            void refreshAllDiagnostics(indexer, diagColl);
        }),
        watcher.onDidCreate(() => {
            indexer.refreshWorkspaceFiles();
            void refreshAllDiagnostics(indexer, diagColl);
        }),
    );

    const selector: vscode.DocumentSelector = { language: 'tl2' };
    context.subscriptions.push(
        vscode.languages.registerDocumentSymbolProvider(selector, new Tl2SymbolProvider(indexer)),
        vscode.languages.registerDefinitionProvider(selector, new Tl2DefinitionProvider(indexer)),
        vscode.languages.registerHoverProvider(selector, new Tl2HoverProvider(indexer)),
        vscode.languages.registerCompletionItemProvider(selector, new Tl2CompletionProvider(indexer)),
    );
}

export function deactivate(): void {
    /* nothing */
}

// ---------- Data model ----------

interface Decl {
    name: string;
    nameRange: vscode.Range;
    fullRange: vscode.Range;
    symbolKind: vscode.SymbolKind;
    detail: string;
    signature: string;
    comment: string;
    children: ConstructorRef[];
    templateParams: string[];
    typeRefs: TypeRefUsage[];
}

interface TypeRefUsage {
    name: string;
    range: vscode.Range;
}

interface ConstructorRef {
    name: string;
    range: vscode.Range;
}

interface FileIndex {
    version: number;
    decls: Decl[];
    byName: Map<string, Decl>;
}

// ---------- Indexer ----------

class Tl2Indexer {
    private cache = new Map<string, FileIndex>();
    private workspaceFilesPromise: Promise<vscode.Uri[]> | null = null;

    indexDocument(doc: vscode.TextDocument): FileIndex {
        const key = doc.uri.toString();
        const existing = this.cache.get(key);
        if (existing && existing.version === doc.version) {
            return existing;
        }
        const decls = parseFile(doc);
        const byName = new Map<string, Decl>();
        for (const d of decls) {
            byName.set(d.name, d);
        }
        const idx: FileIndex = { version: doc.version, decls, byName };
        this.cache.set(key, idx);
        return idx;
    }

    invalidate(uri: vscode.Uri): void {
        this.cache.delete(uri.toString());
    }

    refreshWorkspaceFiles(): void {
        this.workspaceFilesPromise = null;
    }

    private getWorkspaceFiles(): Promise<vscode.Uri[]> {
        if (!this.workspaceFilesPromise) {
            this.workspaceFilesPromise = Promise.resolve(
                vscode.workspace.findFiles('**/*.tl2', '**/node_modules/**', 2000),
            ).then((arr) => Array.from(arr));
        }
        return this.workspaceFilesPromise;
    }

    async ensureWorkspaceIndexed(): Promise<void> {
        const files = await this.getWorkspaceFiles();
        await Promise.all(
            files.map(async (f) => {
                if (this.cache.has(f.toString())) {
                    return;
                }
                try {
                    const doc = await vscode.workspace.openTextDocument(f);
                    this.indexDocument(doc);
                } catch {
                    /* skip */
                }
            }),
        );
    }

    async getAllDeclarations(): Promise<{ uri: vscode.Uri; decl: Decl }[]> {
        await this.ensureWorkspaceIndexed();
        const result: { uri: vscode.Uri; decl: Decl }[] = [];
        for (const [uriStr, idx] of this.cache) {
            const uri = vscode.Uri.parse(uriStr);
            for (const d of idx.decls) {
                result.push({ uri, decl: d });
            }
        }
        return result;
    }

    async getAllNames(): Promise<Set<string>> {
        await this.ensureWorkspaceIndexed();
        const names = new Set<string>();
        for (const idx of this.cache.values()) {
            for (const name of idx.byName.keys()) {
                names.add(name);
            }
        }
        return names;
    }

    async findDeclaration(
        name: string,
        preferUri: vscode.Uri,
    ): Promise<{ uri: vscode.Uri; decl: Decl } | null> {
        const preferKey = preferUri.toString();
        const preferred = this.cache.get(preferKey);
        if (preferred) {
            const d = preferred.byName.get(name);
            if (d) {
                return { uri: preferUri, decl: d };
            }
        }
        await this.ensureWorkspaceIndexed();
        for (const [uriStr, entry] of this.cache) {
            if (uriStr === preferKey) {
                continue;
            }
            const d = entry.byName.get(name);
            if (d) {
                return { uri: vscode.Uri.parse(uriStr), decl: d };
            }
        }
        return null;
    }
}

// ---------- Providers ----------

class Tl2SymbolProvider implements vscode.DocumentSymbolProvider {
    constructor(private indexer: Tl2Indexer) {}

    provideDocumentSymbols(
        document: vscode.TextDocument,
        _token: vscode.CancellationToken,
    ): vscode.DocumentSymbol[] {
        const idx = this.indexer.indexDocument(document);
        return idx.decls.map((d) => {
            const sym = new vscode.DocumentSymbol(
                d.name,
                d.detail,
                d.symbolKind,
                d.fullRange,
                d.nameRange,
            );
            for (const c of d.children) {
                sym.children.push(
                    new vscode.DocumentSymbol(
                        c.name,
                        'constructor',
                        vscode.SymbolKind.EnumMember,
                        c.range,
                        c.range,
                    ),
                );
            }
            return sym;
        });
    }
}

class Tl2DefinitionProvider implements vscode.DefinitionProvider {
    constructor(private indexer: Tl2Indexer) {}

    async provideDefinition(
        document: vscode.TextDocument,
        position: vscode.Position,
        _token: vscode.CancellationToken,
    ): Promise<vscode.Definition | null> {
        this.indexer.indexDocument(document);
        const word = wordAtPosition(document, position);
        if (!word) {
            return null;
        }
        const found = await this.indexer.findDeclaration(word, document.uri);
        if (!found) {
            return null;
        }
        return new vscode.Location(found.uri, found.decl.nameRange);
    }
}

class Tl2HoverProvider implements vscode.HoverProvider {
    constructor(private indexer: Tl2Indexer) {}

    async provideHover(
        document: vscode.TextDocument,
        position: vscode.Position,
        _token: vscode.CancellationToken,
    ): Promise<vscode.Hover | null> {
        this.indexer.indexDocument(document);
        const word = wordAtPosition(document, position);
        if (!word) {
            return null;
        }
        const found = await this.indexer.findDeclaration(word, document.uri);
        if (!found) {
            return null;
        }
        const md = new vscode.MarkdownString();
        md.appendCodeblock(found.decl.signature, 'tl2');
        if (found.decl.comment) {
            md.appendMarkdown('\n' + escapeMarkdown(found.decl.comment));
        }
        const wordRange = document.getWordRangeAtPosition(
            position,
            /[a-zA-Z_][a-zA-Z0-9_]*(?:\.[a-zA-Z_][a-zA-Z0-9_]*)?/,
        );
        return new vscode.Hover(md, wordRange);
    }
}

class Tl2CompletionProvider implements vscode.CompletionItemProvider {
    constructor(private indexer: Tl2Indexer) {}

    async provideCompletionItems(
        _doc: vscode.TextDocument,
        _position: vscode.Position,
        _token: vscode.CancellationToken,
        _ctx: vscode.CompletionContext,
    ): Promise<vscode.CompletionItem[]> {
        const items: vscode.CompletionItem[] = [];

        for (const p of PRIMITIVES) {
            const item = new vscode.CompletionItem(p, vscode.CompletionItemKind.Keyword);
            item.detail = 'built-in primitive';
            item.sortText = `1_${p}`;
            items.push(item);
        }

        const all = await this.indexer.getAllDeclarations();
        const seen = new Set<string>();
        for (const { decl } of all) {
            if (seen.has(decl.name)) {
                continue;
            }
            seen.add(decl.name);
            const item = new vscode.CompletionItem(decl.name, completionKindFor(decl.symbolKind));
            item.detail = decl.detail;
            const md = new vscode.MarkdownString();
            md.appendCodeblock(decl.signature, 'tl2');
            if (decl.comment) {
                md.appendMarkdown('\n' + escapeMarkdown(decl.comment));
            }
            item.documentation = md;
            item.sortText = `2_${decl.name}`;
            items.push(item);
        }

        return items;
    }
}

function completionKindFor(k: vscode.SymbolKind): vscode.CompletionItemKind {
    switch (k) {
        case vscode.SymbolKind.Struct:
            return vscode.CompletionItemKind.Struct;
        case vscode.SymbolKind.Enum:
            return vscode.CompletionItemKind.Enum;
        case vscode.SymbolKind.Interface:
            return vscode.CompletionItemKind.Interface;
        case vscode.SymbolKind.Function:
            return vscode.CompletionItemKind.Function;
        default:
            return vscode.CompletionItemKind.Class;
    }
}

function wordAtPosition(document: vscode.TextDocument, position: vscode.Position): string | null {
    const range = document.getWordRangeAtPosition(
        position,
        /[a-zA-Z_][a-zA-Z0-9_]*(?:\.[a-zA-Z_][a-zA-Z0-9_]*)?/,
    );
    if (!range) {
        return null;
    }
    return document.getText(range);
}

function escapeMarkdown(text: string): string {
    return text.replace(/([\\*_{}\[\]()#+\-!`])/g, '\\$1');
}

// ---------- Diagnostics ----------

async function lintDocument(
    doc: vscode.TextDocument,
    indexer: Tl2Indexer,
    diagColl: vscode.DiagnosticCollection,
): Promise<void> {
    const severity = getDiagnosticSeverity();
    if (severity === null) {
        diagColl.set(doc.uri, []);
        return;
    }

    const fileIdx = indexer.indexDocument(doc);
    const allNames = await indexer.getAllNames();

    const diags: vscode.Diagnostic[] = [];
    for (const decl of fileIdx.decls) {
        const tpSet = new Set(decl.templateParams);
        for (const ref of decl.typeRefs) {
            if (PRIMITIVES.has(ref.name)) {
                continue;
            }
            if (tpSet.has(ref.name)) {
                continue;
            }
            if (allNames.has(ref.name)) {
                continue;
            }
            const d = new vscode.Diagnostic(ref.range, `Unknown type '${ref.name}'`, severity);
            d.source = 'tl2';
            diags.push(d);
        }
    }
    diagColl.set(doc.uri, diags);
}

async function refreshAllDiagnostics(
    indexer: Tl2Indexer,
    diagColl: vscode.DiagnosticCollection,
): Promise<void> {
    for (const doc of vscode.workspace.textDocuments) {
        if (doc.languageId === 'tl2') {
            await lintDocument(doc, indexer, diagColl);
        }
    }
}

function getDiagnosticSeverity(): vscode.DiagnosticSeverity | null {
    const cfg = vscode.workspace.getConfiguration('tl2').get<string>('diagnostics.unknownTypes', 'warning');
    switch (cfg) {
        case 'error':
            return vscode.DiagnosticSeverity.Error;
        case 'warning':
            return vscode.DiagnosticSeverity.Warning;
        case 'information':
            return vscode.DiagnosticSeverity.Information;
        case 'hint':
            return vscode.DiagnosticSeverity.Hint;
        case 'off':
        default:
            return cfg === 'off' ? null : vscode.DiagnosticSeverity.Warning;
    }
}

// ---------- Parser: declarations ----------

function parseFile(document: vscode.TextDocument): Decl[] {
    const originalText = document.getText();
    const stripped = stripComments(originalText);

    const decls: Decl[] = [];
    let i = 0;
    while (i < stripped.length) {
        while (i < stripped.length && /\s/.test(stripped[i])) {
            i++;
        }
        if (i >= stripped.length) {
            break;
        }

        while (stripped[i] === '@') {
            while (i < stripped.length && !/\s/.test(stripped[i])) {
                i++;
            }
            while (i < stripped.length && /\s/.test(stripped[i])) {
                i++;
            }
        }
        if (i >= stripped.length) {
            break;
        }

        const declStart = i;
        let declEnd = stripped.indexOf(';', i);
        if (declEnd < 0) {
            declEnd = stripped.length;
        }

        const body = stripped.substring(declStart, declEnd);
        const decl = parseCombinator(document, body, declStart, declEnd, originalText);
        if (decl) {
            decls.push(decl);
        }

        i = declEnd + 1;
    }
    return decls;
}

function parseCombinator(
    document: vscode.TextDocument,
    body: string,
    absStart: number,
    absEnd: number,
    originalText: string,
): Decl | null {
    let i = 0;
    while (i < body.length && /\s/.test(body[i])) {
        i++;
    }
    if (i >= body.length) {
        return null;
    }

    const nameMatch = /^([a-zA-Z][a-zA-Z0-9_]*)(?:\.([a-zA-Z][a-zA-Z0-9_]*))?/.exec(
        body.substring(i),
    );
    if (!nameMatch) {
        return null;
    }
    const nameOffsetInBody = i;
    const fullName = nameMatch[0];
    i += fullName.length;

    let crc32: string | null = null;
    const crcMatch = /^\s*#([a-fA-F0-9]{1,8})\b/.exec(body.substring(i));
    if (crcMatch) {
        crc32 = crcMatch[1];
        i += crcMatch[0].length;
    }

    while (i < body.length && /\s/.test(body[i])) {
        i++;
    }

    const templateParams: string[] = [];
    if (body[i] === '<' && body[i + 1] !== '=') {
        const tpStart = i + 1;
        let depth = 1;
        i++;
        while (i < body.length && depth > 0) {
            if (body[i] === '<' && body[i + 1] !== '=') {
                depth++;
            } else if (body[i] === '>') {
                depth--;
            }
            i++;
        }
        const tpContent = body.substring(tpStart, i - 1);
        const tpNameRegex = /([a-zA-Z_][a-zA-Z0-9_]*)\s*:\s*(?:Type|#)/g;
        let tm: RegExpExecArray | null;
        while ((tm = tpNameRegex.exec(tpContent)) !== null) {
            templateParams.push(tm[1]);
        }
    }

    while (i < body.length && /\s/.test(body[i])) {
        i++;
    }

    const bodyAfterHeaderStart = i;
    const hasFunctionArrow = containsTopLevelArrow(body, i);
    const hasAlias = body.substring(i, i + 3) === '<=>';
    const hasEq = !hasAlias && body[i] === '=' && body[i + 1] !== '>';

    let symbolKind: vscode.SymbolKind;
    let detail: string;
    let bodyAfterEq = -1;
    let isAlias = false;
    let isFunc = false;

    if (hasFunctionArrow && !hasAlias && !hasEq) {
        symbolKind = vscode.SymbolKind.Function;
        detail = crc32 ? `function #${crc32}` : 'function';
        isFunc = true;
    } else if (hasAlias) {
        symbolKind = vscode.SymbolKind.Interface;
        detail = 'alias';
        isAlias = true;
        i += 3;
    } else if (hasEq) {
        symbolKind = vscode.SymbolKind.Struct;
        detail = crc32 ? `type #${crc32}` : 'type';
        i += 1;
        bodyAfterEq = i;
    } else {
        symbolKind = vscode.SymbolKind.Struct;
        detail = 'type';
    }

    if (!isAlias && !isFunc && bodyAfterEq >= 0) {
        const rest = body.substring(bodyAfterEq);
        if (rest.indexOf('|') >= 0) {
            detail = 'union';
            symbolKind = vscode.SymbolKind.Enum;
        }
    }

    const docTextLen = originalText.length;
    const nameAbsStart = absStart + nameOffsetInBody;
    const nameAbsEnd = nameAbsStart + fullName.length;
    const nameRange = new vscode.Range(
        document.positionAt(nameAbsStart),
        document.positionAt(nameAbsEnd),
    );

    const sigEndAbs = Math.min(absEnd + 1, docTextLen);
    const fullRange = new vscode.Range(
        document.positionAt(nameAbsStart),
        document.positionAt(sigEndAbs),
    );

    const sigFromStripped = body.substring(nameOffsetInBody).replace(/\s+/g, ' ').trim();
    const signature = sigFromStripped + (sigEndAbs > absEnd ? ';' : '');
    const comment = extractLeadingComments(originalText, nameAbsStart);

    const children: ConstructorRef[] = [];
    if (symbolKind === vscode.SymbolKind.Enum && bodyAfterEq >= 0) {
        collectUnionConstructorRanges(document, body, bodyAfterEq, absStart, children);
    }

    const typeRefs: TypeRefUsage[] = [];
    extractDeclTypeRefs(body, bodyAfterHeaderStart, absStart, document, typeRefs);

    return {
        name: fullName,
        nameRange,
        fullRange,
        symbolKind,
        detail,
        signature,
        comment,
        children,
        templateParams,
        typeRefs,
    };
}

function collectUnionConstructorRanges(
    document: vscode.TextDocument,
    body: string,
    bodyAfterEq: number,
    absStart: number,
    out: ConstructorRef[],
): void {
    const segments: { start: number; end: number }[] = [];
    let segStart = bodyAfterEq;
    for (let k = bodyAfterEq; k < body.length; k++) {
        if (body[k] === '|') {
            segments.push({ start: segStart, end: k });
            segStart = k + 1;
        }
    }
    segments.push({ start: segStart, end: body.length });

    for (const seg of segments) {
        const piece = body.substring(seg.start, seg.end);
        const m = /^(\s*)([a-zA-Z][a-zA-Z0-9_]*)/.exec(piece);
        if (!m) {
            continue;
        }
        const after = piece.substring(m[0].length);
        const nextNon = after.replace(/^\s+/, '').charAt(0);
        if (nextNon === ':' || nextNon === '?') {
            continue;
        }
        const ctorName = m[2];
        const ctorAbsStart = absStart + seg.start + m[1].length;
        const ctorAbsEnd = ctorAbsStart + ctorName.length;
        out.push({
            name: ctorName,
            range: new vscode.Range(
                document.positionAt(ctorAbsStart),
                document.positionAt(ctorAbsEnd),
            ),
        });
    }
}

function containsTopLevelArrow(body: string, startIdx: number): boolean {
    let depth = 0;
    for (let k = startIdx; k < body.length - 1; k++) {
        const c = body[k];
        if (c === '<' && body[k + 1] !== '=') {
            depth++;
        } else if (c === '>') {
            if (depth > 0) {
                depth--;
            }
        } else if (c === '[') {
            depth++;
        } else if (c === ']') {
            if (depth > 0) {
                depth--;
            }
        } else if (depth === 0 && c === '=' && body[k + 1] === '>') {
            return true;
        } else if (depth === 0 && c === '=' && body[k + 1] !== '>' && body[k - 1] !== '<') {
            return false;
        }
    }
    return false;
}

function extractLeadingComments(text: string, declStart: number): string {
    let lineStart = declStart;
    while (lineStart > 0 && text[lineStart - 1] !== '\n') {
        lineStart--;
    }

    const lines: string[] = [];
    let cursor = lineStart - 1;
    while (cursor >= 0) {
        let curLineStart = cursor;
        while (curLineStart > 0 && text[curLineStart - 1] !== '\n') {
            curLineStart--;
        }
        const line = text.substring(curLineStart, cursor);
        const trimmed = line.trim();

        if (trimmed === '') {
            break;
        }
        if (trimmed.startsWith('//')) {
            const stripped = trimmed.replace(/^\/\/\s?/, '');
            lines.unshift(stripped);
            cursor = curLineStart - 1;
        } else {
            break;
        }
    }
    return lines.join('\n');
}

function stripComments(text: string): string {
    let result = '';
    let i = 0;
    while (i < text.length) {
        if (text[i] === '/' && text[i + 1] === '/') {
            while (i < text.length && text[i] !== '\n') {
                result += ' ';
                i++;
            }
        } else if (text[i] === '/' && text[i + 1] === '*') {
            result += '  ';
            i += 2;
            while (i + 1 < text.length && !(text[i] === '*' && text[i + 1] === '/')) {
                result += text[i] === '\n' ? '\n' : ' ';
                i++;
            }
            if (i + 1 < text.length) {
                result += '  ';
                i += 2;
            } else {
                while (i < text.length) {
                    result += ' ';
                    i++;
                }
            }
        } else {
            result += text[i];
            i++;
        }
    }
    return result;
}

// ---------- Parser: type references ----------

function extractDeclTypeRefs(
    body: string,
    startInBody: number,
    absStart: number,
    document: vscode.TextDocument,
    out: TypeRefUsage[],
): void {
    let i = startInBody;
    while (i < body.length && /\s/.test(body[i])) {
        i++;
    }
    if (i >= body.length) {
        return;
    }

    if (body.substring(i, i + 3) === '<=>') {
        scanTypeRefChunk(body, i + 3, body.length, absStart, document, out);
        return;
    }

    if (body[i] === '=' && body[i + 1] !== '>') {
        scanStructOrUnion(body, i + 1, body.length, absStart, document, out);
        return;
    }

    const arrowIdx = findTopLevelArrowInRange(body, i, body.length);
    if (arrowIdx >= 0) {
        scanFields(body, i, arrowIdx, absStart, document, out);
        scanStructOrUnion(body, arrowIdx + 2, body.length, absStart, document, out);
    }
}

function scanStructOrUnion(
    body: string,
    start: number,
    end: number,
    absStart: number,
    doc: vscode.TextDocument,
    out: TypeRefUsage[],
): void {
    const pipes = findTopLevelPipesInRange(body, start, end);
    if (pipes.length === 0) {
        if (hasTopLevelColonInRange(body, start, end)) {
            scanFields(body, start, end, absStart, doc, out);
        } else {
            scanTypeRefChunk(body, start, end, absStart, doc, out);
        }
        return;
    }
    const segments: { s: number; e: number }[] = [];
    let prev = start;
    for (const p of pipes) {
        segments.push({ s: prev, e: p });
        prev = p + 1;
    }
    segments.push({ s: prev, e: end });
    for (const seg of segments) {
        scanUnionSegment(body, seg.s, seg.e, absStart, doc, out);
    }
}

function scanUnionSegment(
    body: string,
    start: number,
    end: number,
    absStart: number,
    doc: vscode.TextDocument,
    out: TypeRefUsage[],
): void {
    let i = start;
    while (i < end && /\s/.test(body[i])) {
        i++;
    }
    if (i >= end) {
        return;
    }
    const rem = body.substring(i, end);
    const m = /^([a-zA-Z][a-zA-Z0-9_]*)(?:\.([a-zA-Z][a-zA-Z0-9_]*))?/.exec(rem);
    if (!m) {
        return;
    }
    let f = i + m[0].length;
    while (f < end && /\s/.test(body[f])) {
        f++;
    }
    const isField = body[f] === ':' || (body[f] === '?' && body[f + 1] === ':');
    if (isField) {
        scanFields(body, i, end, absStart, doc, out);
        return;
    }
    const after = i + m[0].length;
    if (hasTopLevelColonInRange(body, after, end)) {
        scanFields(body, after, end, absStart, doc, out);
    } else {
        scanTypeRefChunk(body, after, end, absStart, doc, out);
    }
}

function scanFields(
    body: string,
    start: number,
    end: number,
    absStart: number,
    doc: vscode.TextDocument,
    out: TypeRefUsage[],
): void {
    let i = start;
    while (i < end) {
        while (i < end && /\s/.test(body[i])) {
            i++;
        }
        if (i >= end) {
            break;
        }
        const rem = body.substring(i, end);
        const fm = /^([a-zA-Z][a-zA-Z0-9_]*|_)(\??)\s*(:)/.exec(rem);
        if (!fm) {
            i++;
            continue;
        }
        i += fm[0].length;
        const typeEnd = findFieldEnd(body, i, end);
        scanTypeRefChunk(body, i, typeEnd, absStart, doc, out);
        i = typeEnd;
    }
}

function findFieldEnd(body: string, start: number, end: number): number {
    let i = start;
    let angleDepth = 0;
    let squareDepth = 0;
    while (i < end) {
        const c = body[i];
        if (c === '<' && body[i + 1] !== '=') {
            angleDepth++;
        } else if (c === '>' && angleDepth > 0) {
            angleDepth--;
        } else if (c === '[') {
            squareDepth++;
        } else if (c === ']' && squareDepth > 0) {
            squareDepth--;
        } else if (angleDepth === 0 && squareDepth === 0) {
            if (c === '|' || c === ';') {
                return i;
            }
            if (c === '=' && body[i + 1] === '>') {
                return i;
            }
            if (matchesFieldNameAt(body, i, end)) {
                return i;
            }
        }
        i++;
    }
    return i;
}

function matchesFieldNameAt(body: string, pos: number, end: number): boolean {
    if (pos >= end) {
        return false;
    }
    const rem = body.substring(pos, end);
    return /^([a-zA-Z][a-zA-Z0-9_]*|_)(\??)\s*:/.test(rem);
}

function scanTypeRefChunk(
    body: string,
    start: number,
    end: number,
    absStart: number,
    doc: vscode.TextDocument,
    out: TypeRefUsage[],
): void {
    const text = body.substring(start, end);
    const identRegex = /([a-zA-Z][a-zA-Z0-9_]*)(?:\.([a-zA-Z][a-zA-Z0-9_]*))?/g;
    let m: RegExpExecArray | null;
    while ((m = identRegex.exec(text)) !== null) {
        const name = m[0];
        if (name === 'Type') {
            continue;
        }
        const localIdx = m.index;
        const absStartPos = absStart + start + localIdx;
        out.push({
            name,
            range: new vscode.Range(
                doc.positionAt(absStartPos),
                doc.positionAt(absStartPos + name.length),
            ),
        });
    }
}

function findTopLevelPipesInRange(body: string, start: number, end: number): number[] {
    const positions: number[] = [];
    let angleDepth = 0;
    let squareDepth = 0;
    for (let i = start; i < end; i++) {
        const c = body[i];
        if (c === '<' && body[i + 1] !== '=') {
            angleDepth++;
        } else if (c === '>' && angleDepth > 0) {
            angleDepth--;
        } else if (c === '[') {
            squareDepth++;
        } else if (c === ']' && squareDepth > 0) {
            squareDepth--;
        } else if (angleDepth === 0 && squareDepth === 0 && c === '|') {
            positions.push(i);
        }
    }
    return positions;
}

function hasTopLevelColonInRange(body: string, start: number, end: number): boolean {
    let angleDepth = 0;
    let squareDepth = 0;
    for (let i = start; i < end; i++) {
        const c = body[i];
        if (c === '<' && body[i + 1] !== '=') {
            angleDepth++;
        } else if (c === '>' && angleDepth > 0) {
            angleDepth--;
        } else if (c === '[') {
            squareDepth++;
        } else if (c === ']' && squareDepth > 0) {
            squareDepth--;
        } else if (angleDepth === 0 && squareDepth === 0 && c === ':') {
            return true;
        }
    }
    return false;
}

function findTopLevelArrowInRange(body: string, start: number, end: number): number {
    let angleDepth = 0;
    let squareDepth = 0;
    for (let i = start; i < end - 1; i++) {
        if (body.substring(i, i + 3) === '<=>') {
            i += 2;
            continue;
        }
        const c = body[i];
        if (c === '<' && body[i + 1] !== '=') {
            angleDepth++;
        } else if (c === '>' && angleDepth > 0) {
            angleDepth--;
        } else if (c === '[') {
            squareDepth++;
        } else if (c === ']' && squareDepth > 0) {
            squareDepth--;
        } else if (angleDepth === 0 && squareDepth === 0 && c === '=' && body[i + 1] === '>') {
            return i;
        }
    }
    return -1;
}
