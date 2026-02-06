// Please add -types- to the end of section if section ends with functions

//----- Test Multiple native wrappers

// innerType {t:Type} a:t = InnerType t;

// allBoxed a:(innerType (Maybe int)) b:(innerType (Maybe Int))
// c:(innerType (Maybe int)) d:(innerType (Maybe Int)) = AllBoxed;

// Multiline comment before combinator
//----- Test Multiple native wrappers
//
//----- Test Multiple bools
boolFalse#bc799737 = Bool;
boolTrue#997275b5 = Bool;

pair {X:Type} {Y:Type} a:X b:Y = Pair X Y;

//----- Test Common Types

//resultFalse#27930a7b {t:Type} = Maybe t;
//resultTrue#3f9c8ef8 {t:Type} t = Maybe t;

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
