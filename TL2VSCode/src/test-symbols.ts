/// <reference types="node" />
// Standalone smoke test for parser logic (no vscode runtime).
// Usage:
//   npx ts-node src/test-symbols.ts <file.tl2>          # show symbols + refs
//   npx ts-node src/test-symbols.ts --lint <dir>        # workspace lint (unknown types)

import * as fs from 'fs';
import * as path from 'path';

const PRIMITIVES = new Set([
    'int32', 'uint32', 'int64', 'uint64', 'float32', 'float64',
    'byte', 'string', 'bool', 'bit', 'true',
]);

interface TypeRef { name: string; offset: number }
interface Decl {
    name: string;
    kind: string;
    detail: string;
    signature: string;
    comment: string;
    children: string[];
    templateParams: string[];
    typeRefs: TypeRef[];
}

// ----- parser -----

function stripComments(text: string): string {
    let result = '';
    let i = 0;
    while (i < text.length) {
        if (text[i] === '/' && text[i + 1] === '/') {
            while (i < text.length && text[i] !== '\n') { result += ' '; i++; }
        } else if (text[i] === '/' && text[i + 1] === '*') {
            result += '  '; i += 2;
            while (i + 1 < text.length && !(text[i] === '*' && text[i + 1] === '/')) {
                result += text[i] === '\n' ? '\n' : ' '; i++;
            }
            if (i + 1 < text.length) { result += '  '; i += 2; }
            else while (i < text.length) { result += ' '; i++; }
        } else { result += text[i]; i++; }
    }
    return result;
}

function containsTopLevelArrow(body: string, startIdx: number): boolean {
    let depth = 0;
    for (let k = startIdx; k < body.length - 1; k++) {
        const c = body[k];
        if (c === '<' && body[k + 1] !== '=') depth++;
        else if (c === '>') { if (depth > 0) depth--; }
        else if (c === '[') depth++;
        else if (c === ']') { if (depth > 0) depth--; }
        else if (depth === 0 && c === '=' && body[k + 1] === '>') return true;
        else if (depth === 0 && c === '=' && body[k + 1] !== '>' && body[k - 1] !== '<') return false;
    }
    return false;
}

function extractLeadingComments(text: string, declStart: number): string {
    let lineStart = declStart;
    while (lineStart > 0 && text[lineStart - 1] !== '\n') lineStart--;
    const lines: string[] = [];
    let cursor = lineStart - 1;
    while (cursor >= 0) {
        let curLineStart = cursor;
        while (curLineStart > 0 && text[curLineStart - 1] !== '\n') curLineStart--;
        const trimmed = text.substring(curLineStart, cursor).trim();
        if (trimmed === '') break;
        if (trimmed.startsWith('//')) {
            lines.unshift(trimmed.replace(/^\/\/\s?/, ''));
            cursor = curLineStart - 1;
        } else break;
    }
    return lines.join('\n');
}

function findTopLevelArrowInRange(body: string, start: number, end: number): number {
    let a = 0, s = 0;
    for (let i = start; i < end - 1; i++) {
        if (body.substring(i, i + 3) === '<=>') { i += 2; continue; }
        const c = body[i];
        if (c === '<' && body[i + 1] !== '=') a++;
        else if (c === '>' && a > 0) a--;
        else if (c === '[') s++;
        else if (c === ']' && s > 0) s--;
        else if (a === 0 && s === 0 && c === '=' && body[i + 1] === '>') return i;
    }
    return -1;
}

function findTopLevelPipesInRange(body: string, start: number, end: number): number[] {
    const out: number[] = [];
    let a = 0, s = 0;
    for (let i = start; i < end; i++) {
        const c = body[i];
        if (c === '<' && body[i + 1] !== '=') a++;
        else if (c === '>' && a > 0) a--;
        else if (c === '[') s++;
        else if (c === ']' && s > 0) s--;
        else if (a === 0 && s === 0 && c === '|') out.push(i);
    }
    return out;
}

function hasTopLevelColonInRange(body: string, start: number, end: number): boolean {
    let a = 0, s = 0;
    for (let i = start; i < end; i++) {
        const c = body[i];
        if (c === '<' && body[i + 1] !== '=') a++;
        else if (c === '>' && a > 0) a--;
        else if (c === '[') s++;
        else if (c === ']' && s > 0) s--;
        else if (a === 0 && s === 0 && c === ':') return true;
    }
    return false;
}

function matchesFieldNameAt(body: string, pos: number, end: number): boolean {
    if (pos >= end) return false;
    return /^([a-zA-Z][a-zA-Z0-9_]*|_)(\??)\s*:/.test(body.substring(pos, end));
}

function findFieldEnd(body: string, start: number, end: number): number {
    let i = start, a = 0, s = 0;
    while (i < end) {
        const c = body[i];
        if (c === '<' && body[i + 1] !== '=') a++;
        else if (c === '>' && a > 0) a--;
        else if (c === '[') s++;
        else if (c === ']' && s > 0) s--;
        else if (a === 0 && s === 0) {
            if (c === '|' || c === ';') return i;
            if (c === '=' && body[i + 1] === '>') return i;
            if (matchesFieldNameAt(body, i, end)) return i;
        }
        i++;
    }
    return i;
}

function scanTypeRefChunk(body: string, start: number, end: number, absStart: number, out: TypeRef[]) {
    const text = body.substring(start, end);
    const re = /([a-zA-Z][a-zA-Z0-9_]*)(?:\.([a-zA-Z][a-zA-Z0-9_]*))?/g;
    let m;
    while ((m = re.exec(text)) !== null) {
        if (m[0] === 'Type') continue;
        out.push({ name: m[0], offset: absStart + start + m.index });
    }
}

function scanFields(body: string, start: number, end: number, absStart: number, out: TypeRef[]) {
    let i = start;
    while (i < end) {
        while (i < end && /\s/.test(body[i])) i++;
        if (i >= end) break;
        const fm = /^([a-zA-Z][a-zA-Z0-9_]*|_)(\??)\s*(:)/.exec(body.substring(i, end));
        if (!fm) { i++; continue; }
        i += fm[0].length;
        const te = findFieldEnd(body, i, end);
        scanTypeRefChunk(body, i, te, absStart, out);
        i = te;
    }
}

function scanUnionSegment(body: string, start: number, end: number, absStart: number, out: TypeRef[]) {
    let i = start;
    while (i < end && /\s/.test(body[i])) i++;
    if (i >= end) return;
    const m = /^([a-zA-Z][a-zA-Z0-9_]*)(?:\.([a-zA-Z][a-zA-Z0-9_]*))?/.exec(body.substring(i, end));
    if (!m) return;
    let f = i + m[0].length;
    while (f < end && /\s/.test(body[f])) f++;
    const isField = body[f] === ':' || (body[f] === '?' && body[f + 1] === ':');
    if (isField) { scanFields(body, i, end, absStart, out); return; }
    const after = i + m[0].length;
    if (hasTopLevelColonInRange(body, after, end)) scanFields(body, after, end, absStart, out);
    else scanTypeRefChunk(body, after, end, absStart, out);
}

function scanStructOrUnion(body: string, start: number, end: number, absStart: number, out: TypeRef[]) {
    const pipes = findTopLevelPipesInRange(body, start, end);
    if (pipes.length === 0) {
        if (hasTopLevelColonInRange(body, start, end)) scanFields(body, start, end, absStart, out);
        else scanTypeRefChunk(body, start, end, absStart, out);
        return;
    }
    const segs: { s: number; e: number }[] = [];
    let prev = start;
    for (const p of pipes) { segs.push({ s: prev, e: p }); prev = p + 1; }
    segs.push({ s: prev, e: end });
    for (const seg of segs) scanUnionSegment(body, seg.s, seg.e, absStart, out);
}

function extractDeclTypeRefs(body: string, startInBody: number, absStart: number, out: TypeRef[]) {
    let i = startInBody;
    while (i < body.length && /\s/.test(body[i])) i++;
    if (i >= body.length) return;
    if (body.substring(i, i + 3) === '<=>') { scanTypeRefChunk(body, i + 3, body.length, absStart, out); return; }
    if (body[i] === '=' && body[i + 1] !== '>') { scanStructOrUnion(body, i + 1, body.length, absStart, out); return; }
    const arrowIdx = findTopLevelArrowInRange(body, i, body.length);
    if (arrowIdx >= 0) {
        scanFields(body, i, arrowIdx, absStart, out);
        scanStructOrUnion(body, arrowIdx + 2, body.length, absStart, out);
    }
}

function parseFile(text: string): Decl[] {
    const stripped = stripComments(text);
    const decls: Decl[] = [];
    let i = 0;
    while (i < stripped.length) {
        while (i < stripped.length && /\s/.test(stripped[i])) i++;
        if (i >= stripped.length) break;
        while (stripped[i] === '@') {
            while (i < stripped.length && !/\s/.test(stripped[i])) i++;
            while (i < stripped.length && /\s/.test(stripped[i])) i++;
        }
        if (i >= stripped.length) break;
        const ds = i;
        let de = stripped.indexOf(';', i);
        if (de < 0) de = stripped.length;
        const body = stripped.substring(ds, de);
        const decl = parseCombinator(body, ds, de, text);
        if (decl) decls.push(decl);
        i = de + 1;
    }
    return decls;
}

function parseCombinator(body: string, absStart: number, absEnd: number, originalText: string): Decl | null {
    let i = 0;
    while (i < body.length && /\s/.test(body[i])) i++;
    if (i >= body.length) return null;
    const nm = /^([a-zA-Z][a-zA-Z0-9_]*)(?:\.([a-zA-Z][a-zA-Z0-9_]*))?/.exec(body.substring(i));
    if (!nm) return null;
    const fullName = nm[0];
    const nameOffsetInBody = i;
    i += fullName.length;

    let crc32: string | null = null;
    const cm = /^\s*#([a-fA-F0-9]{1,8})\b/.exec(body.substring(i));
    if (cm) { crc32 = cm[1]; i += cm[0].length; }
    while (i < body.length && /\s/.test(body[i])) i++;

    const templateParams: string[] = [];
    if (body[i] === '<' && body[i + 1] !== '=') {
        const tpStart = i + 1;
        let depth = 1; i++;
        while (i < body.length && depth > 0) {
            if (body[i] === '<' && body[i + 1] !== '=') depth++;
            else if (body[i] === '>') depth--;
            i++;
        }
        const tpContent = body.substring(tpStart, i - 1);
        const tpRe = /([a-zA-Z_][a-zA-Z0-9_]*)\s*:\s*(?:Type|#)/g;
        let tm;
        while ((tm = tpRe.exec(tpContent)) !== null) templateParams.push(tm[1]);
    }
    while (i < body.length && /\s/.test(body[i])) i++;

    const bodyAfterHeader = i;
    const hasFn = containsTopLevelArrow(body, i);
    const hasAlias = body.substring(i, i + 3) === '<=>';
    const hasEq = !hasAlias && body[i] === '=' && body[i + 1] !== '>';

    let kind = 'Struct', detail = 'type', bodyAfterEq = -1, isAlias = false, isFunc = false;
    if (hasFn && !hasAlias && !hasEq) { kind = 'Function'; detail = crc32 ? `function #${crc32}` : 'function'; isFunc = true; }
    else if (hasAlias) { kind = 'Interface'; detail = 'alias'; isAlias = true; }
    else if (hasEq) { kind = 'Struct'; detail = crc32 ? `type #${crc32}` : 'type'; i += 1; bodyAfterEq = i; }

    if (!isAlias && !isFunc && bodyAfterEq >= 0) {
        if (body.substring(bodyAfterEq).indexOf('|') >= 0) { detail = 'union'; kind = 'Enum'; }
    }

    const nameAbsStart = absStart + nameOffsetInBody;
    const sigEnd = Math.min(absEnd + 1, originalText.length);
    const sigFromStripped = body.substring(nameOffsetInBody).replace(/\s+/g, ' ').trim();
    const signature = sigFromStripped + (sigEnd > absEnd ? ';' : '');
    const comment = extractLeadingComments(originalText, nameAbsStart);

    const children: string[] = [];
    if (kind === 'Enum' && bodyAfterEq >= 0) {
        const segs: { start: number; end: number }[] = [];
        let segStart = bodyAfterEq;
        for (let k = bodyAfterEq; k < body.length; k++) {
            if (body[k] === '|') { segs.push({ start: segStart, end: k }); segStart = k + 1; }
        }
        segs.push({ start: segStart, end: body.length });
        for (const seg of segs) {
            const piece = body.substring(seg.start, seg.end);
            const m = /^(\s*)([a-zA-Z][a-zA-Z0-9_]*)/.exec(piece);
            if (!m) continue;
            const after = piece.substring(m[0].length);
            const next = after.replace(/^\s+/, '').charAt(0);
            if (next === ':' || next === '?') continue;
            children.push(m[2]);
        }
    }

    const typeRefs: TypeRef[] = [];
    extractDeclTypeRefs(body, bodyAfterHeader, absStart, typeRefs);

    return { name: fullName, kind, detail, signature, comment, children, templateParams, typeRefs };
}

// ----- modes -----

function showFile(file: string) {
    const text = fs.readFileSync(file, 'utf-8');
    const decls = parseFile(text);
    console.log(`File: ${path.basename(file)} — ${decls.length} top-level decls\n`);
    for (const d of decls) {
        console.log(`[${d.kind}] ${d.name} (${d.detail})`);
        if (d.templateParams.length) console.log(`    tpl: <${d.templateParams.join(', ')}>`);
        if (d.signature.length > 100) console.log(`    sig: ${d.signature.substring(0, 97)}...`);
        else console.log(`    sig: ${d.signature}`);
        if (d.comment) console.log(`    doc: ${d.comment.replace(/\n/g, ' | ')}`);
        for (const c of d.children) console.log(`    └ ctor ${c}`);
        if (d.typeRefs.length) {
            const refNames = d.typeRefs.map((r) => r.name).join(' ');
            console.log(`    refs: ${refNames}`);
        }
    }
}

function lintDir(dir: string) {
    const files: string[] = [];
    function walk(d: string) {
        for (const e of fs.readdirSync(d)) {
            const full = path.join(d, e);
            const st = fs.statSync(full);
            if (st.isDirectory()) walk(full);
            else if (e.endsWith('.tl2')) files.push(full);
        }
    }
    walk(dir);

    console.log(`Scanning ${files.length} .tl2 files in ${dir}\n`);
    const allNames = new Set<string>();
    const fileDecls = new Map<string, Decl[]>();
    for (const f of files) {
        const decls = parseFile(fs.readFileSync(f, 'utf-8'));
        fileDecls.set(f, decls);
        for (const d of decls) allNames.add(d.name);
    }

    let problems = 0;
    for (const [file, decls] of fileDecls) {
        for (const decl of decls) {
            const tpSet = new Set(decl.templateParams);
            for (const ref of decl.typeRefs) {
                if (PRIMITIVES.has(ref.name)) continue;
                if (tpSet.has(ref.name)) continue;
                if (allNames.has(ref.name)) continue;
                console.log(`${path.relative(dir, file)}: in ${decl.name} — unknown type '${ref.name}' @ offset ${ref.offset}`);
                problems++;
            }
        }
    }
    console.log(`\n${problems} unknown-type reference(s) total.`);
}

// ----- main -----

const args = process.argv.slice(2);
if (args[0] === '--lint' && args[1]) lintDir(args[1]);
else if (args[0]) showFile(args[0]);
else { console.error('usage: ts-node test-symbols.ts <file.tl2>'); console.error('   or: ts-node test-symbols.ts --lint <dir>'); process.exit(1); }
