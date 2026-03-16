boolFalse#bc799737 = Bool;
boolTrue#997275b5 = Bool;

// BASE TYPES
int#a8509bda ? = Int;
long#22076cba ? = Long;

string#b5286e24 ? = String;

resultFalse#27930a7b {t:Type} = Maybe t;
resultTrue#3f9c8ef8 {t:Type} t = Maybe t;

pair {X:Type} {Y:Type} a:X b:Y = Pair X Y;

vector#1cb5c415 {t:Type} # [t] = Vector t;
tuple#9770768a {t:Type} {n:#} [t] = Tuple t n;

dictionaryField {t:Type} key:string value:t = DictionaryField t;
dictionary#1f4c618f {t:Type} %(Vector %(DictionaryField t)) = Dictionary t;

//dictionaryAnyField {k:Type} {v:Type} key:k value:v = DictionaryAnyField k v;
//dictionaryAny#1f4c6190 {k:Type} {v:Type} # [(dictionaryAnyField k v)] = DictionaryAny k v;

true = True; // this can be used as void type and serialized to empty array in PHP

map {X:Type} {Y:Type} key:X value:Y = Map X Y;

//a.inner {I:#} a:I*[int] = a.Inner I;
//a.middle {W:#} {P:Type} {Q:Type} a:(a.inner W) b:(pair P Q) = a.Middle W P Q;
//a.top2 n:# m:# c:(a.middle 4 (pair (a.inner 5) (a.inner n)) (a.Inner 3)) = a.Top2;

//@read a.call1#12345678 x:int => Int;

//@read a.call2#12345679 x:int => True;

// curl.myHren curl.myHren = curl.MyHren;

@readwrite tree_stats.preferMaster {X:Type} query:!X => X;
@any rpcProxy.diagonal {X:Type} query:!X => Vector (Maybe X);
@any rpcProxy.diagonalTargets {X:Type} offset:int limit:int query:!X => Vector (%Vector (Maybe X));
