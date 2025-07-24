// BASE TYPES
int#a8509bda ? = Int;
long#22076cba ? = Long;

int32 int = Int32;
int64 long = Int64;
string#b5286e24 ? = String;

vector#1cb5c415 {t:Type} # [t] = Vector t;
tuple#9770768a {t:Type} {n:#} [t] = Tuple t n;

dictionaryField {t:Type} key:string value:t = DictionaryField t;
dictionary#1f4c618f {t:Type} %(Vector %(DictionaryField t)) = Dictionary t;
dictionaryFieldAny {k:Type} {v:Type} key:k value:v = DictionaryFieldAny k v;
dictionaryAny#1f4c6190 {k:Type} {v:Type} # [(dictionaryFieldAny k v)] = DictionaryAny k v;

true = True;

resultFalse#27930a7b {t:Type} = Maybe t;
resultTrue#3f9c8ef8 {t:Type} t = Maybe t;

pair {X:Type} {Y:Type} x:X y:Y = Pair X Y;

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
cases.testDictAny dict:(dictionaryAny double int) = cases.TestDictAny;

// BLOCK: TEST UNIONS
cases.testMaybe value:(Maybe int) = cases.TestMaybe;

cases.testUnion1 value:int = cases.TestUnion;
cases.testUnion2 value:string = cases.TestUnion;

cases.testUnionContainer value:cases.TestUnion = cases.TestUnionContainer;

cases.testEnum1 = cases.TestEnum;
cases.testEnum2 = cases.TestEnum;
cases.testEnum3 = cases.TestEnum;

cases.testEnumContainer value:cases.TestEnum = cases.TestEnumContainer;

// BLOCK: TEST FIELD MASKS
cases.testRecursiveFieldMask f0:# f1:f0.0?# f2:f1.1?# t1:f0.0?%True t2:f1.1?%True t3:f2.2?%True = cases.TestRecursiveFieldmask;
cases.testLocalFieldmask f1:# f2:f1.0?# f3:f2.1?%True f4:f2.1?%True = cases.TestLocalFieldmask;

cases.testOutFieldMask {f:#} f1:f.0?# f2:f.3?%True f3:f*[int] = cases.TestOutFieldMask f;
cases.testOutFieldMaskContainer f:# inner:(cases.testOutFieldMask f) = cases.TestOutFieldMaskContainer;
cases.testBeforeReadBitValidation n:# a:n.0?n*[int] b:n.1?n*[int] = cases.TestBeforeReadBitValidation;

cases.testAllPossibleFieldConfigs {outer:#} local:#
    f00:int                 // no fieldmask, no nat dependencies in type
    f01:%True               // no fieldmask, true type
    f02:local*[int]         // no fieldmask, have some local nat dependencies in type
    f03:outer*[int]         // no fieldmask, have only external nat dependencies in type
    f10:local.0?int         // local fieldmask, no nat dependencies in type
    f11:local.1?%True       // local fieldmask, true type
    f12:local.2?local*[int] // local fieldmask, have some local nat dependencies in type
    f13:local.3?outer*[int] // local fieldmask, have only external nat dependencies in type
    f20:outer.0?int         // outer fieldmask, no nat dependencies in type
    f21:outer.1?%True       // outer fieldmask, true type
    f22:outer.2?local*[int] // outer fieldmask, have some local nat dependencies in type
    f23:outer.3?outer*[int] // outer fieldmask, have only external nat dependencies in type
= cases.TestAllPossibleFieldConfigs outer;

cases.testAllPossibleFieldConfigsContainer outer:#
    value:(cases.testAllPossibleFieldConfigs outer)
= cases.TestAllPossibleFieldConfigsContainer;

// BLOCK: TEST INPLACING ARGUMENTS
cases.testInplaceStructArgs a1:# a2:# a3:# arg:cases.inplace1<a1, a2, a3, int> = cases.TestInplaceStructArgs;
cases.testInplaceStructArgs2 a1:# a2:# a3:# arg:cases.inplace1<a1, a2, a3, pair<tuple<int, a3>, tuple<int, a2>>> = cases.TestInplaceStructArgs2;

cases.inplace1 {a1:#} {a2:#} {a3:#} {X:Type} value:cases.inplace2<a2, a3, a1, X> = cases.Inplace1 a1 a2 a3 X;
cases.inplace2 {a1:#} {a2:#} {a3:#} {X:Type} value:cases.inplace3<a2, a3, a1, tuple<X, 2>> = cases.Inplace2 a1 a2 a3 X;
cases.inplace3 {a1:#} {a2:#} {a3:#} {X:Type} value:pair<tuple<X, a2>, tuple<X, a3>> = cases.Inplace3 a1 a2 a3 X;

// BLOCK: Bytes vesrion
cases_bytes.testArray n:# arr:n*[string] = cases_bytes.TestArray;
cases_bytes.testVector arr:(vector string) = cases_bytes.TestVector;
cases_bytes.testTuple tpl:(tuple string 4) = cases_bytes.TestTuple;

// will be map
cases_bytes.testDictStringString dict:(dictionary string) = cases_bytes.TestDictStringString;
// will be map
cases_bytes.testDictString dict:(dictionary int) = cases_bytes.TestDictString;
// will be map
cases_bytes.testDictInt dict:(dictionaryAny int int) = cases_bytes.TestDictInt;
// will be array of tuples
cases_bytes.testDictAny dict:(dictionaryAny double int) = cases_bytes.TestDictAny;

cases_bytes.testEnum1 = cases_bytes.TestEnum;
cases_bytes.testEnum2 = cases_bytes.TestEnum;
cases_bytes.testEnum3 = cases_bytes.TestEnum;
cases_bytes.testEnumContainer value:cases.TestEnum = cases_bytes.TestEnumContainer;


// BLOCK: BENCHMARKS
benchmarks.vruhash#d31bd0fd low:long high:long = benchmarks.VruHash; // We store xxh3 components in LE endianness, as common in TL

benchmarks.vruposition#32792c04 fields_mask:#
    commit_bit:         fields_mask.0?true // we set this bit only if block with this bit set is committed, so we have no more than 1 such uncommitted block.
    meta_block:         fields_mask.1?true // this block contains meta data of VR protocol
    split_payload:      fields_mask.3?true // this position splits payload. If this set, meta_block and commit_bit are both not set
    rotation_block:     fields_mask.5?true // If set, binlog should rotate and this block should be in new file
    canonical_hash:     fields_mask.15?true // If set, hash is stored in canonical byte order.
    payload_offset:     long // Motivation for storing it in each block - when random seeking into the middle of file, we must be sure without scanning from beginning
    block_time_nano:    long // Only for debug, not used in any logic
    hash:               benchmarks.vruhash
    file_offset:        long
    seq_number:         fields_mask.14?long // incremented after each block.
= benchmarks.VruPosition;

benchmarks.vrutoytopLevelUnionBig
    next_positions:(vector benchmarks.vruposition)
= benchmarks.VrutoyTopLevelUnion;

benchmarks.vrutoytopLevelUnionEmpty
= benchmarks.VrutoyTopLevelUnion;

benchmarks.vrutoyTopLevelContainer
    value: benchmarks.VrutoyTopLevelUnion
= benchmarks.VrutoyTopLevelContainer;

benchmarks.vrutoyPositions {n: #}
    next_positions:n*[benchmarks.vruposition]
= benchmarks.VrutoyPositions n;

benchmarks.vrutoyTopLevelContainerWithDependency
    n:# value:(benchmarks.vrutoyPositions n)
= benchmarks.VrutoyTopLevelContainerWithDependency;

// Boolean emulation
boolFalse#bc799737 = Bool;
boolTrue#997275b5 = Bool;

casesTL2.testObject
    n:#
    f1:n.0?true
    f2:Bool
    f3:vector<Bool>
    f4:(benchmarks.vrutoyPositions n)
    f5:n.1?Bool
    f6:vector<benchmarks.VrutoyTopLevelUnion>
    f7:n.14?vector<True>
    = casesTL2.TestObject;

casesTL2.testVectorBool x:vector<Bool> = casesTL2.TestVectorBool;
casesTL2.testArrayFixedBool
    a1:1*[Bool]
    a2:0*[Bool]
    a3:7*[Bool]
    a4:8*[Bool]
    a5:11*[Bool]
    a6:16*[Bool] = casesTL2.TestArrayFixedBool;

casesTL2.testArrayFlexibleBool
    n:#
    arr:n*[Bool] = casesTL2.TestArrayFlexibleBool;

casesTL2.testFixedParam
    x: casesTL2.testObjectWithParam<4> = casesTL2.TestFixedParam;

casesTL2.testObjectWithParam {n:#}
    x:n.0?int
    y:n*[int] = casesTL2.TestObjectWithParam n;

casesTL2.testParamsGeneration
    n1:#
    n2:#
    x1: casesTL2.testObjectWithMuiltiParams<4, 1>
    x2: casesTL2.testObjectWithMuiltiParams<n1, 2>
    x3: casesTL2.testObjectWithMuiltiParams<3, n2>
    x4: casesTL2.testObjectWithMuiltiParams<n1, n2>
    = casesTL2.TestParamsGeneration;

casesTL2.testObjectWithMuiltiParams {n:#} {m:#}
    f1:n.0?n*[int]
    f2:m.0?m*[int] = casesTL2.TestObjectWithMuiltiParams n m;


@read casesTL2.testFunctionNoDep1 x:int => Int;
@read casesTL2.testFunctionNoDep2 x:int => True;
@read casesTL2.testFunctionNoDep3 x:int => Vector<int>;
@read casesTL2.testFunctionNoDep4 x:int => Bool;
@read casesTL2.testFunctionNoDep5 x:int => casesTL2.TestObject;

@read casesTL2.testFunctionWithDep1 n:# => Tuple<int, n>;
@read casesTL2.testFunctionWithDep2 n:# => benchmarks.VrutoyPositions n;
