// Please add -types- to the end of section if section ends with functions

//----- Test Multiple native wrappers
int#a8509bda ? = Int;
long#22076cba ? = Long;

// innerType {t:Type} a:t = InnerType t;

// allBoxed a:(innerType (Maybe int)) b:(innerType (Maybe Int))
// c:(innerType (Maybe int)) d:(innerType (Maybe Int)) = AllBoxed;

// Multiline comment before combinator
//----- Test Multiple native wrappers
//
//----- Test Multiple bools
boolFalse#bc799737 = Bool;
boolTrue#997275b5 = Bool;

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

dictionaryAnyField {k:Type} {v:Type} key:k value:v = DictionaryAnyField k v;
dictionaryAny#1f4c6190 {k:Type} {v:Type} # [(dictionaryAnyField k v)] = DictionaryAny k v;

intKeyDictionaryField {t:Type}
    key:int
    value:t
    = IntKeyDictionaryField t;

intKeyDictionary#07bafc42 {t:Type}
    %(Vector %(intKeyDictionaryField t))
    = IntKeyDictionary t;

longKeyDictionaryField {t:Type}
    key:long
    value:t
    = LongKeyDictionaryField t;

longKeyDictionary#b424d8f1 {t:Type}
    %(Vector %(longKeyDictionaryField t))
    = LongKeyDictionary t;

ab.useDictUgly#fb9ce817 n:#
    //c:dictionaryAny<pair<int,int>, int>
    d:intKeyDictionary<tuple<string,n>>
    e:longKeyDictionary<tuple<string,n>>
    f:dictionary<pair<tuple<string,n>, tuple<string,n>>>
    g:Dictionary<pair<tuple<string,n>, tuple<string,n>>>
    x:dictionaryAny<int, pair<int,int>>
    y:dictionaryAny<long, pair<int,int>>
    z:dictionaryAny<string, pair<int,int>>
     = ab.UseDictUgly;

ab.useDictString#3325d884 fields_mask:#
  tags:    %(Dictionary string)
  = ab.UseDictString;
