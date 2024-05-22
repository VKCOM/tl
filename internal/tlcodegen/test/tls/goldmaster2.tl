---functions---
@read call1 x:ab.typeA = ab.TypeB;
@read call2 x:ab.typeA = cd.TypeB;
@read call3 x:ab.typeA = TypeB;
@read call4 x:cd.typeA = ab.TypeB;
@read call5 x:cd.typeA = cd.TypeB;
@read call6 x:cd.typeA = TypeB;
@read call7 x:typeA = ab.TypeB;
@read call8 x:typeA = cd.TypeB;
@read call9 x:typeA = TypeB;

@readwrite service5.insert#7cf362ba flags:# persistent:flags.0?%True = service5.Output;
@readwrite service5Long.insert#7cf362bb flags:# persistent:flags.0?%True = service5Long.Output;

---types---
