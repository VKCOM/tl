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
//----- Test Common Types

//resultFalse#27930a7b {t:Type} = Maybe t;
//resultTrue#3f9c8ef8 {t:Type} t = Maybe t;

ab.innerX {n:#} a:n*[int] = ab.InnerX n;

ab.useInnerX1 n:# a:ab.innerX<n> = ab.UseInnerX1;
ab.useInnerX2 n:# a:ab.innerX<2> = ab.UseInnerX2;
ab.useInnerX3 n:# a:ab.innerX<3> = ab.UseInnerX3;
