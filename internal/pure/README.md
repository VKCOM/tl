# hybrid TL1 + TL2 pure kernel

This kernel combines relatively simple TL2 kernel with complicated TL1 kernel.

Primitive types are difficult, because they have different names in TL1/TL2.

And TL1 types are themselves complicated, because they have 2 names - TypeDecl.Name and Construct.Name, `ab.t1 = cd.T2`, 
but only if they are not unions, which use Construct.Name only for naming data structure for variant.

Also, TL1 type references can be bare or boxed (explicitly with % sign, or by using lower or upper case).

So, we inveneted those rules:

Kernel has `types` map element for each possible type reference.
For primitives `int32` and `int` are added initially, later if boxed wrapper is declared, `Int` is also added.

During type resolution, we do not in general normalize, only substitute template arguments. So `%int` could remain `%int`.

When we want type canonical string, we always look up each name in `types`, then normalize reference.

TL2 types normalize into themselves, because TL2 has no boxed references.

Primitives, like `int32`, `int`, `Int` with or without percent sign always normalize to `+int32` (boxed) or `int32`.

Unions always normalize to `+Union` with explicit boxed sign (TODO - try to run kernel without it).

Bare and boxed references to single-constructor type `ab.con = ab.Type;` normalize to `ab.con` or `+ab.con`.

Then, when language generator wants to construct global names, it can select some strategy to use.

We also explicitly add legacy names to some primitives (`int`, `nat` and `long`), so we generate
VectorInt not VectorInt32 in go generator. Later, we'll replace those to TL2 names.

## TODO

TL1 template argument/field normalized name collision

modifiers in new kernel

## Migration

plan

### fields masks - local, external

### JSON format - simple types, unions