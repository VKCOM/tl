
pair {X:Type} {Y:Type} a:X b:Y = Pair X Y;

a.inner {I:#} a:I*[int] = a.Inner I;
a.middle {W:#} {P:Type} {Q:Type} a:(a.inner W) b:(pair P Q) = a.Middle W P Q;
a.top2 n:# m:# c:(a.middle m (pair (a.inner n) (a.inner n)) (a.Inner 3)) d:(a.Inner 4) = a.Top2;
a.top3 n:# m:# c:(a.middle 5 (a.inner n) (a.Inner 3)) = a.Top3;
