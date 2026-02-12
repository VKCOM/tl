//boolFalse#bc799737 = Bool;
//boolTrue#997275b5 = Bool;
// BASE TYPES
int#a8509bda ? = Int;

vector#1cb5c415 {t:Type} # [t] = Vector t;
tuple#9770768a {t:Type} {n:#} [t] = Tuple t n;

dictionaryField {t:Type} key:string value:t = DictionaryField t;
dictionary#1f4c618f {t:Type} %(Vector %(DictionaryField t)) = Dictionary t;
dictionaryAnyField {k:Type} {v:Type} key:k value:v = DictionaryAnyField k v;
dictionaryAny#1f4c6190 {k:Type} {v:Type} # [(dictionaryAnyField k v)] = DictionaryAny k v;


 cases.testPrimitive
     a:int
     b:Int
     c:%Int
     d:vector<int>
     e:vector<Int>
     f:vector<%Int>
     g:vector<Int>
     h:Vector<Int>
     i:%Vector<Int>
     j:vector<#>
 = cases.TestPrimitive;

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

//pair {X:Type} {Y:Type} a:X b:Y = Pair X Y;

//a.inner {I:#} a:I*[int] = a.Inner I;
//a.middle {W:#} {P:Type} {Q:Type} a:(a.inner W) b:(pair P Q) = a.Middle W P Q;
//a.top2 n:# m:# c:(a.middle m (pair (a.inner n) (a.inner n)) (a.Inner 3)) = a.Top2;


//resultFalse#27930a7b {t:Type} = Maybe t;
//resultTrue#3f9c8ef8 {t:Type} t = Maybe t;

//maybeWrapper {T:Type} {N:#} a:(Maybe (tuple T N)) = MaybeWrapper T N;
//testMaybe n:# a:(Maybe int) b:(Maybe Int)
//    c:(Maybe (vector int)) d:(Maybe (tuple int 3)) // comment for d field
//    e:(Maybe (tuple int n)) f:(maybeWrapper int 3) = TestMaybe;
