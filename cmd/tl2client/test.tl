boolFalse#bc799737 = Bool;
boolTrue#997275b5 = Bool;
// BASE TYPES
int#a8509bda ? = Int;

vector#1cb5c415 {t:Type} # [t] = Vector t;
tuple#9770768a {t:Type} {n:#} [t] = Tuple t n;

//dictionaryField {t:Type} key:string value:t = DictionaryField t;
//dictionary#1f4c618f {t:Type} %(Vector %(DictionaryField t)) = Dictionary t;
//dictionaryAnyField {k:Type} {v:Type} key:k value:v = DictionaryAnyField k v;
//dictionaryAny#1f4c6190 {k:Type} {v:Type} # [(dictionaryAnyField k v)] = DictionaryAny k v;


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
 = cases.TestPrimitive;

// will be map
// cases.testDictString dict:(dictionary int) = cases.TestDictString;
// will be map
// cases.testDictInt dict:(dictionaryAny int int) = cases.TestDictInt;
// will be array of tuples
// cases.testDictAny dict:(dictionaryAny long int) = cases.TestDictAny;

// cases.testAllDicts
//    f1:cases.testDictString
//    f2:cases.testDictInt
//    f3:cases.testDictAny
//    = cases.TestAllDicts;
