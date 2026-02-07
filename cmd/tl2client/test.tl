boolFalse#bc799737 = Bool;
boolTrue#997275b5 = Bool;
// BASE TYPES
int#a8509bda ? = Int;
long#22076cba ? = Long;
string#b5286e24 ? = String;

vector#1cb5c415 {t:Type} # [t] = Vector t;
tuple#9770768a {t:Type} {n:#} [t] = Tuple t n;

dictionaryField {t:Type} key:string value:t = DictionaryField t;
dictionary#1f4c618f {t:Type} %(Vector %(DictionaryField t)) = Dictionary t;
dictionaryAnyField {k:Type} {v:Type} key:k value:v = DictionaryAnyField k v;
dictionaryAny#1f4c6190 {k:Type} {v:Type} # [(dictionaryAnyField k v)] = DictionaryAny k v;

true = True;

resultFalse#27930a7b {t:Type} = Maybe t;
resultTrue#3f9c8ef8 {t:Type} t = Maybe t;

pair {X:Type} {Y:Type} x:X y:Y = Pair X Y;

ab.innerX {n:#} a:n*[int] = ab.InnerX n;

ab.useInnerX1 n:# a:%ab.InnerX<n> = ab.UseInnerX1;
ab.useInnerX2 n:# a:ab.InnerX<2> = ab.UseInnerX2;
ab.useInnerX3 n:# a:ab.innerX<3> = ab.UseInnerX3;

ab.innerY {n:#} a:n*[int] = ab.InnerY n;

ab.useInnerY2 n:# a:ab.InnerY<2> = ab.UseInnerY2;
ab.useInnerY3 n:# a:ab.innerY<3> = ab.UseInnerY3;

a.inner {I:#} a:I*[int] = a.Inner I;
a.middle {W:#} {P:Type} {Q:Type} a:(a.inner W) b:(pair P Q) = a.Middle W P Q;
a.top2 n:# m:# c:(a.middle m (pair (a.inner n) (a.inner n)) (a.Inner 3)) = a.Top2;

ch_proxy.emptyOutput#ff8f7db8 = ch_proxy.Output;
ch_proxy.stringOutput#dc170ff4 http_code:int response:string = ch_proxy.Output;

// BLOCK: TEST RECURSIVE TYPES
cases.myCycle1 fields_mask:# a:fields_mask.0?cases.myCycle2 = cases.MyCycle1;
cases.myCycle2 fields_mask:# a:fields_mask.0?cases.myCycle3 = cases.MyCycle2;
cases.myCycle3 fields_mask:# a:fields_mask.0?cases.myCycle1 = cases.MyCycle3;

// BLOCK: TEST BRACKETS (arrays, vectors, dicitinaries etc.)
cases.testArray n:# arr:n*[int] = cases.TestArray;
cases.testVector arr:(vector int) = cases.TestVector;
cases.testTuple tpl:(tuple int 4) = cases.TestTuple;

// will be map
cases.testDictString dict:(dictionary int) = cases.TestDictString;
// will be map
cases.testDictInt dict:(dictionaryAny int int) = cases.TestDictInt;
// will be array of tuples
cases.testDictAny dict:(dictionaryAny long int) = cases.TestDictAny;

cases.testAllDicts
    f1:cases.testDictString
    f2:cases.testDictInt
    f3:cases.testDictAny
    = cases.TestAllDicts;

// BLOCK: TEST UNIONS
cases.testMaybe value:(Maybe int) = cases.TestMaybe;

cases.testUnion1 value:int = cases.TestUnion;
cases.testUnion2 value:string = cases.TestUnion;
cases.testUnion3 int = cases.TestUnion;
cases.testUnion4 string = cases.TestUnion;
// cases.testUnion5 bool = cases.TestUnion;
// cases.testUnion6 bit = cases.TestUnion;

cases.testUnionContainer value:cases.TestUnion = cases.TestUnionContainer;

cases.testEnum1 = cases.TestEnum;
cases.testEnum2 = cases.TestEnum;
cases.testEnum3 = cases.TestEnum;

cases.testEnumContainer value:cases.TestEnum = cases.TestEnumContainer;

// BLOCK: TEST FIELD MASKS
cases.testRecursiveFieldMask f0:# f1:f0.0?# f2:f1.1?# t1:f0.0?%True t2:f1.1?%True t3:f2.2?%True = cases.TestRecursiveFieldmask;
cases.testLocalFieldmask f1:# f2:f1.0?# f3:f2.1?%True f4:f2.1?%True = cases.TestLocalFieldmask;


cases.testOutFieldMask {f:#} {f_s:#} f1:f.0?# f2:f.3?%True f3:f_s*[int] = cases.TestOutFieldMask f f_s;
cases.testOutFieldMaskContainer f:# f_s:# inner:(cases.testOutFieldMask f f_s) = cases.TestOutFieldMaskContainer;
cases.testBeforeReadBitValidation n:# n_s:# a:n.0?n_s*[int] b:n.1?n_s*[int] = cases.TestBeforeReadBitValidation;

cases.testAllPossibleFieldConfigs {outer:#} {outer_s:#} local:# local_s:#
    f00:int                 // no fieldmask, no nat dependencies in type
    f01:%True               // no fieldmask, true type
    f02:local_s*[int]         // no fieldmask, have some local nat dependencies in type
    f03:outer_s*[int]         // no fieldmask, have only external nat dependencies in type
    f10:local.0?int         // local fieldmask, no nat dependencies in type
    f11:local.1?%True       // local fieldmask, true type
    f12:local.2?local_s*[int] // local fieldmask, have some local nat dependencies in type
    f13:local.3?outer_s*[int] // local fieldmask, have only external nat dependencies in type
    f20:outer.0?int         // outer fieldmask, no nat dependencies in type
    f21:outer.1?%True       // outer fieldmask, true type
    f22:outer.2?local_s*[int] // outer fieldmask, have some local nat dependencies in type
    f23:outer.3?outer_s*[int] // outer fieldmask, have only external nat dependencies in type
= cases.TestAllPossibleFieldConfigs outer outer_s;

cases.testAllPossibleFieldConfigsContainer outer:# outer_s:#
    value:(cases.testAllPossibleFieldConfigs outer outer_s)
= cases.TestAllPossibleFieldConfigsContainer;
