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

---types---

videoAdsService.fullPostID#4dda7b2f
    fields_mask:              #
    owner_id:                 long
    post_id:                  long
    bit1:           fields_mask.2?true
    bool1:          fields_mask.3?Bool
    = videoAdsService.FullPostID;

videoAdsService.checkParams#6742a1bb videoAdsService.fullPostID = videoAdsService.CheckParams;

videoAdsService.checkFields#cf0cd638
    fields_mask:                  #
    can_show_banner:              fields_mask.0?true
    ads_disabled:                 fields_mask.1?true
    legal_partner:                fields_mask.2?videoAdsService.checkParams
    ads_type:                     fields_mask.3?videoAdsService.CheckParams
    = videoAdsService.CheckFields;

@read
videoAdsService.getPromoConfigs#a8bdd93d
    arg:videoAdsService.checkParams
    => videoAdsService.CheckParams;
