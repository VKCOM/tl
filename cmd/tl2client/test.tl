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

dictionaryFieldAny {k:Type} {v:Type} key:k value:v = DictionaryFieldAny k v;
dictionaryAny#1f4c6190 {k:Type} {v:Type} # [(dictionaryFieldAny k v)] = DictionaryAny k v;

true = True; // this can be used as void type and serialized to empty array in PHP

// testX a:# c:a*[int] b:a.0?int = TestX;
// testY a:# c:tuple<int, a> b:a.0?int = TestY;

a.inner {I:#} a:I*[int] = a.Inner I;
a.middle {W:#} {P:Type} {Q:Type} a:(a.inner W) b:(pair P Q) = a.Middle W P Q;
a.top2 n:# m:# c:(a.middle m (pair (a.inner n) (a.inner n)) (a.Inner 3)) = a.Top2;

curl.responseOk#57d7a33a
  field_mask:   #
  http_code:    int
  http_headers: %(Dictionary string)
  http_body:    field_mask.0?string
  = curl.Response;

curl.responseError#af514651
  field_mask:   #
  error_code:   int
  error_string: string
  = curl.Response;

---functions---

@kphp
curl.request#3f5a4651
  field_mask: #
  method:             string
  url:                string
  headers:            field_mask.0?%(Dictionary string)
  body:               field_mask.1?string
  timeout:            field_mask.2?int
  timeout_connection: field_mask.3?int
  max_redirs:         field_mask.4?int
  ssl_version:        field_mask.5?int
  resolve_ipv:        field_mask.6?int
  http_proxy:         field_mask.7?string
  limit:              field_mask.8?int
  route:              field_mask.9?string
  version:            field_mask.10?int
  origin:             field_mask.11?string
  tls_cert:           field_mask.12?string
  tls_cert_pk:        field_mask.13?string
  = curl.Response;