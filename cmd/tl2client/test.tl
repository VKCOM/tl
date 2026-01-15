int#a8509bda ? = Int;
//long#22076cba ? = Long;
//string#b5286e24 ? = String;

//myInt32#ba59e151 int32s = MyInt32;
//myInt64#1d95db9d int64s = MyInt64;

// innerType {t:Type} a:t = InnerType t;

// allBoxed a:(innerType (Maybe int)) b:(innerType (Maybe Int))
// c:(innerType (Maybe int)) d:(innerType (Maybe Int)) = AllBoxed;

// Multiline comment before combinator
//----- Test Multiple native wrappers
//
//multiPoint#0e1ae81e a:3*[int] b:3*[Int] c:3*[int32s] d:3*[Int32s] e:3*[myInt32] f:3*[MyInt32]= MultiPoint;

//nativeWrappers#344ddf50 a:int b:Int c:int32s d:Int32s e:myInt32 f:MyInt32 = NativeWrappers; // combinator comment
//----- Test Multiple bools
//boolFalse#bc799737 = Bool;
//boolTrue#997275b5 = Bool;

//a.myTrue#00000001 = a.MyBool;
//a.myFalse#00000002 = a.MyBool;
//----- Test Common Types

//resultFalse#27930a7b {t:Type} = Maybe t;
//resultTrue#3f9c8ef8 {t:Type} t = Maybe t;

//pair {X:Type} {Y:Type} a:X b:Y = Pair X Y;

vector#1cb5c415 {t:Type} # [t] = Vector t;
tuple#9770768a {t:Type} {n:#} [t] = Tuple t n;

//pair {X:Type} {Y:Type} a:X b:Y = Pair X Y;

testVectorTuple n:#
    a:vector<int> b:Vector<int>
    c:vector<Int> d:Vector<Int>
    e:tuple<int, n> f:Tuple<int, n>
    g:tuple<Int, n> h:Tuple<Int, n>
    i:tuple<int, 3> j:Tuple<int, 3>
    k:tuple<Int, 3> l:Tuple<Int, 3>
    = TestVectorTuple;

//---- test bytes version propagation
//noStr x:int = NoStr;
//useStr x:string = UseStr;
//halfStr x:noStr y:useStr = HalfStr;
//ab.topLevel1 a:noStr b:noStr = ab.TopLevel1;

//ab.topLevel2 a:noStr b:useStr = ab.TopLevel2;
// skipped comment

// comment for TopLevel3
// also comment for TopLevel3
//cd.topLevel3 a:useStr b:halfStr = cd.TopLevel3;
//----- Test Dependent types
//usePair pa:pair<uint32, uint32> = UsePair;

//a.inner {I:#} a:int = a.Inner I;
//a.middle {W:#} {P:Type} {Q:Type} a:a.inner<W> b:Pair<P, Q> = a.Middle W P Q;
//a.top2 n:# m:# c:a.middle<m, pair<a.inner<n>, a.inner<n>>, a.Inner<3>> = a.Top2;

// a.middle<P1, pair<a.inner<P2>, a.inner<P3>>, a.Inner<3>>(w, pxi, pyi)
// W:P1 P:pair<a.inner<P2>, a.inner<P3>> Q:a.Inner<3>
