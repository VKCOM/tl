pair {X:Type} {Y:Type} a:X b:Y = Pair X Y;

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

a.inner {I:#} a:int = a.Inner I;
a.middle {W:#} {P:Type} {Q:Type} a:a.inner<W> b:pair<P, Q> = a.Middle W P Q;
a.top2 n:# m:# c:a.middle<m, pair<a.inner<n>, a.inner<n>>, a.Inner<3>> = a.Top2;

// a.middle<P1, pair<a.inner<P2>, a.inner<P3>>, a.Inner<3>>(w, pxi, pyi)
// W:P1 P:pair<a.inner<P2>, a.inner<P3>> Q:a.Inner<3>
