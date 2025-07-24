// Please add -types- to the end of section if section ends with functions

//----- Test Multiple native wrappers
int#a8509bda ? = Int;
long#22076cba ? = Long;

int32 int = Int32;
int64 long = Int64;
string#b5286e24 ? = String;

myInt32 int32 = MyInt32;
myInt64 int64 = MyInt64;

// innerType {t:Type} a:t = InnerType t;

// allBoxed a:(innerType (Maybe int)) b:(innerType (Maybe Int))
// c:(innerType (Maybe int)) d:(innerType (Maybe Int)) = AllBoxed;

// Multiline comment before combinator
//----- Test Multiple native wrappers
//
multiPoint a:3*[int] b:3*[Int] c:3*[int32] d:3*[Int32] e:3*[myInt32] f:3*[MyInt32]= MultiPoint;

nativeWrappers a:int b:Int c:int32 d:Int32 e:myInt32 f:MyInt32 = NativeWrappers; // combinator comment
//----- Test Multiple bools
boolFalse#bc799737 = Bool;
boolTrue#997275b5 = Bool;

a.myTrue#00000001 = a.MyBool;
a.myFalse#00000002 = a.MyBool;
//----- Test Common Types

resultFalse#27930a7b {t:Type} = Maybe t;
resultTrue#3f9c8ef8 {t:Type} t = Maybe t;

pair {X:Type} {Y:Type} a:X b:Y = Pair X Y;

vector#1cb5c415 {t:Type} # [t] = Vector t;
tuple#9770768a {t:Type} {n:#} [t] = Tuple t n;

dictionaryField {t:Type} key:string value:t = DictionaryField t;
dictionary#1f4c618f {t:Type} %(Vector %(DictionaryField t)) = Dictionary t;

// dictionaryField {t:Type} key:string value:t = DictionaryField t;
// dictionary#1f4c618f {t:Type} # [(dictionaryField t)] = Dictionary t;

dictionaryFieldAny {k:Type} {v:Type} key:k value:v = DictionaryFieldAny k v;
dictionaryAny#1f4c6190 {k:Type} {v:Type} # [(dictionaryFieldAny k v)] = DictionaryAny k v;

true = True; // this can be used as void type and serialized to empty array in PHP
useTrue#dfdd4180 fm:# a:fm.0?true b:fm.1?True c:true d:True e:fm.2?Bool = UseTrue; // we give warnings for many fields of this kind

//---- test recursive struct without union as result
list {X:Type} flag:# head:flag.0?X tail:flag.0?(list X) = List X;

//---- test bytes version propagation
noStr x:int = NoStr;
useStr x:string = UseStr;
halfStr x:noStr y:useStr = HalfStr;
ab.topLevel1 a:noStr b:noStr = ab.TopLevel1;

ab.topLevel2 a:noStr b:useStr = ab.TopLevel2;
// skipped comment

// comment for TopLevel3
// also comment for TopLevel3
cd.topLevel3 a:useStr b:halfStr = cd.TopLevel3;
//----- Test Dependent types
a.inner {I:#} a:I*[int] = a.Inner I;
a.middle {W:#} {P:Type} {Q:Type} a:(a.inner W) b:(pair P Q) = a.Middle W P Q;
a.top2 n:# m:# c:(a.middle m (pair (a.inner n) (a.inner n)) (a.Inner 3)) = a.Top2;
//---- test maybe
maybeWrapper {T:Type} {N:#} a:(Maybe (tuple T N)) = MaybeWrapper T N;
testMaybe n:# a:(Maybe int) b:(Maybe Int)
    c:(Maybe (vector int)) d:(Maybe (tuple int 3)) // comment for d field
    e:(Maybe (tuple int n)) f:(maybeWrapper int 3) = TestMaybe;

ab.myType x:int = ab.MyType;
cd.myType x:int = cd.MyType;

testMaybe2 n:# a:(Maybe int) g:(Maybe ab.myType) = TestMaybe2;
ab.testMaybe n:# a:(Maybe int) b:(Maybe ab.myType) c:(Maybe cd.myType) d:(Maybe ab.MyType) = ab.TestMaybe;

innerMaybe {X:#} a:(Maybe (tuple int X)) = InnerMaybe X;
maybeTest1 n:# a:(Maybe int) b:(Maybe Int) c:(Maybe (tuple int 3)) d:(Maybe (tuple Int 3))

// This comment is before field e
// Second line of comment before field e
    e:(InnerMaybe n) f:(InnerMaybe 0) g:(Maybe (vector int)) h:(Maybe (Vector int))
    i:(Maybe (Tuple int 3)) j:(Maybe (Tuple Int 0)) = MaybeTest1;
//----- Test Recursive types
myNat fields_mask:# a:fields_mask.0?myNat = MyNat;

myZero = MyNat2;
myPlus a:MyNat2 = MyNat2; // struct with field

myZero3 = MyNat3;
myPlus3 MyNat3 = MyNat3; // typedef

// ban such things
cycleTuple n:# a:n.0?(Tuple cycleTuple 2) b:(Tuple cycleTuple n) c:n.2?3*[int] = CycleTuple;

hren next:(Maybe hren) = Hren;
//----- Test Recursive between namespaces
cyc1.myCycle fields_mask:# a:fields_mask.0?cyc2.myCycle = cyc1.MyCycle;
cyc2.myCycle fields_mask:# a:(Maybe cyc3.myCycle) b:a.Color = cyc2.MyCycle;
cyc3.myCycle fields_mask:# a:(vector cyc1.myCycle) = cyc3.MyCycle;

ab.useCycle a:cyc2.myCycle b:(vector a.Color) = ab.UseCycle;
cd.useCycle a:cyc2.myCycle b:(Maybe a.Color) = cd.UseCycle;
//----- Test Bracket replace
inner {X:#} a:int = Inner X;
replace1 {n:#} a:[int] = Replace1 n;
replace2 n:# a:[int] = Replace2;
replace3 a:3*[int] = Replace3;
replace4 {n:#} a:n*[int] = Replace4 n;
replace5 n:# a:n*[int] = Replace5;
replace6 # a:[int] = Replace6;
replace7 n:# m:# a:n*[m*[int]] = Replace7;
replace8 n:# a:[3*[int]] = Replace8;
replace9 n:# a:[3*[(Inner n)]] = Replace9;
replace10 n:# a:[3*[a:int b:int]] = Replace10;
replace11 {n:#} {X:Type} k:# a:k*[b:n.0?3*[(Pair int X)] c:n*[int]] = Replace11 n X;
replace12 n:# a:[3*[a:n.0?int]] = Replace12;
replace13 {n:#} {X:Type} k:# a:k*[b:n.0?k*[(Pair int X)] c:k.0?n*[int]] = Replace13 n X;
replace14 {n:#} {X:Type} k:# a:k*[b:n.0?3*[x:int y:k*[X]] c:int] = Replace14 n X;
replace15 n:# a:n*[x:int y:int] b:n*[x:long y:long] c:n*[x:int y:int z:int]= Replace15;
// replace16 n:int k:# m:# a:z*[b:# c:z*[int] k:# d:k*[int]] = Replace16; - interesting errors
// replace16 n:int k:# m:# a:[b:# c:n*[int] k:# d:k*[int]] = Replace16; - interesting errors
replace17 # x:[int] a:int # y:[int] = Replace17;
replace18 # a:[# [# [int]]] = Replace18;

replace n:# a:(replace1 n) a1:(replace1 3) b:replace2
 c:replace3 d:(replace4 n) d1:(replace4 3) e:replace5 g:replace6 h:replace7
 i:replace8 j:replace9 k:replace10 l:(replace11 n long) m:replace12 o:(replace13 n long)
 p:(replace14 n long) q:replace15
= Replace;


//----- Unions
unionArgsXXX1 {X:Type} {Y:#} x:(tuple X Y) = UnionArgsXXX X Y;
unionArgsXXX2 {X:Type} {Y:#} a:int = UnionArgsXXX X Y;

unionArgsUse k:# n:# a:(UnionArgsXXX int k) b:(UnionArgsXXX long n) = UnionArgsUse;

a.uNionA a:int = a.MyUnion;
au.nionA b:int = a.MyUnion;

ab.empty = ab.Response; // Union comment - Empty
ab.code x:int = ab.Response; // Union comment - Code
ab.alias int = ab.Response;

// Also union comment
// Multiline Alias
cd.response x:int str:string = ab.Response;


useResponse n:# x:[ab.Response] = UseResponse;

// remove common prefix
ab.counterChangeRequestPeriodsMany
  objects_peridos:(vector int)
  = ab.CounterChangeRequestPeriods;

ab.counterChangeRequestPeriodsOne
  period:long
  = ab.CounterChangeRequestPeriods;
//----- enums
a.color = a.Color;
a.red   = a.Color; // Enum comment - red
a.green = a.Color; // Enum comment - green
b.red   = a.Color;
a.blue  = a.Color; // Enum comment - blue, but
// continues on the next line

//----- Test Dictionaries
usefulService.userEntityPaymentItemPromo#24c7ec9f
  {fields_mask: #}
  content: string
  = usefulService.UserEntityPaymentItemPromo fields_mask;

usefulService.userEntityPaymentItem#4f798680
  {fields_mask: #}
  id: string
  promo: (Maybe (usefulService.UserEntityPaymentItemPromo fields_mask))
  = usefulService.UserEntityPaymentItem fields_mask;

usefulService.getUserEntityResult#e494698a
  {fields_mask: #}
  entity_id: string
  payment_items: (Maybe (Dictionary (usefulService.UserEntityPaymentItem fields_mask)))
  = usefulService.GetUserEntityResult fields_mask;

@readwrite
usefulService.getUserEntity#3c857e52
  fields_mask: #
  stage_id: fields_mask.0?string
  => Maybe (usefulService.GetUserEntityResult fields_mask);

---functions---

---types---

dictionaryElemStrange {v:Type} key:# value:key.31?v = DictionaryElemStrange v;
dictionaryElemUgly {f:#} {k:Type} {v:Type} key:f.0?k value:f.1?v = DictionaryElemUgly f k v;
dictionaryElem {k:Type} {v:Type} key:k value:v = DictionaryElem k v;

useDictUgly n:#
    # a:[(dictionaryElemUgly n int string)]
    # b:[(dictionaryElemStrange string)]
    # c:[(dictionaryElem (pair int int) int)]
    # d:[(dictionaryElem (tuple string n) int)]
    # e:[(dictionaryElem (pair Bool a.Color) int)]
    # f:[(dictionaryElem (pair float double) int)]
    # g:[(dictionaryElem (pair int (pair multiPoint string)) int)]
    # x:[(dictionaryElem int (pair int int))]
    # y:[(dictionaryElem long (pair int int))]
    # z:[(dictionaryElem string (pair int int))]
     = UseDictUgly;

ab.useDictString#3325d884 fields_mask:#
  tags:    %(Dictionary string)
  = ab.UseDictString;
//----- Test function calls to and from namespaces
typeA x:int = TypeA;
typeB x:int = TypeB;
typeC x:ab.typeA = TypeC;
typeD x:cd.typeA = TypeD;

ab.typeA x:int = ab.TypeA;
ab.typeB x:ab.typeA = ab.TypeB;
ab.typeC x:cd.typeA = ab.TypeC;
ab.typeD x:typeA = ab.TypeD;

cd.typeA x:int = cd.TypeA;
cd.typeB x:ab.typeA = cd.TypeB;
cd.typeC x:ab.typeA = cd.TypeC;
cd.typeD x:typeA = cd.TypeD;

service5.emptyOutput#ff8f7db8 = service5.Output;
service5.stringOutput#dc170ff4 http_code:int response:string = service5.Output;

service5Long.emptyOutput#ff8f7db9 = service5Long.Output;
service5Long.stringOutput#dc170ff5 http_code:long response:string = service5Long.Output;

//hrenAlias {T:Type} {N:#} (Tuple T N) = HrenAlias T N;
//topHrenAlias our_n:# our_m:# a:our_n*[int] b:(Tuple int our_m) c:our_n*[our_m*[int]] => HrenAlias (Inner our_n) our_m;
//topHrenAlias2 our_n:# our_m:# => HrenAlias (Tuple (Inner 3) our_n) our_m;
//topHrenAlias3 our_n:# our_m:# => Vector (Vector (Inner 3));

@read ab.call1 x:ab.typeA => ab.TypeB;
@read ab.call2 x:ab.typeA => cd.TypeB;
@read ab.call3 x:ab.typeA => TypeB;
@read ab.call4 x:cd.typeA => ab.TypeB;
@read ab.call5 x:cd.typeA => cd.TypeB;
@read ab.call6 x:cd.typeA => TypeB;
@read ab.call7 x:typeA => ab.TypeB;
@read ab.call8 x:typeA => cd.TypeB;
@read ab.call9 x:typeA => TypeB;
@read ab.call10 => a.Color;
@read ab.call11 x:a.Color => a.Color;

// @write logs2.addIndex      type_id:int field:%String => Bool; // Test that %String is canonically resolved into primitive string
