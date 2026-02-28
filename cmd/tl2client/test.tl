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

useDictUgly#fb9ce817 n:#
    // c:dictionaryAny<pair<int,int>, int>
    d:intKeyDictionary<tuple<string,n>>
    e:longKeyDictionary<tuple<string,n>>
    // f:dictionary<pair<tuple<string,n>, tuple<string,n>>>
//    x:dictionaryAny<int, pair<int,int>>
//    y:dictionaryAny<long, pair<int,int>>
//    z:dictionaryAny<string, pair<int,int>>
     = UseDictUgly;

//videoAdsService.getLegalPartnerAdsSettingsManyResponseOk#4db649d8
//    fields_mask:              #
//    settings:                 vector<videoAdsService.legalPartnerAdsSettings>
//    = videoAdsService.GetLegalPartnerAdsSettingsManyResponse;
//
//videoAdsService.getLegalPartnerAdsSettingsManyResponseError#143f3aaa
//    fields_mask:              #
//    error_code:               int
//    error:                    string
//    = videoAdsService.GetLegalPartnerAdsSettingsManyResponse;
//
//videoAdsService.legalPartnerAdsSettingsVideoRules#ee6ba4fa
//    fields_mask:              #
//    video_id:                 int
//    //rules:                    videoAdsService.legalPartnerAdsSettingsRules
//    = videoAdsService.LegalPartnerAdsSettingsVideoRules;
//
//
//videoAdsService.legalPartnerAdsSettings#59ee3756
//    fields_mask:              #
//    //common_rules:             fields_mask.0?videoAdsService.legalPartnerAdsSettingsRules
//    video_rules:              fields_mask.1?vector<videoAdsService.legalPartnerAdsSettingsVideoRules>
//    //playlist_rules:           fields_mask.2?vector<videoAdsService.legalPartnerAdsSettingsPlaylistRules>
//    = videoAdsService.LegalPartnerAdsSettings;
//
//
////nonMigrating.example
////    use:               vector<videoAdsService.legalPartnerAdsSettings>
////    use2:              videoAdsService.LegalPartnerAdsSettings
////    = nonMigrating.Example;
////
////nonMigrating2.example
////    use:               videoAdsService.LegalPartnerAdsSettings
////    use2:              vector<videoAdsService.legalPartnerAdsSettings>
////    = nonMigrating2.Example;
//
//---functions---
//
//@read
//videoAdsService.getLegalPartnerAdsSettingsMany#98505e85
//    fields_mask:              #
////    requests:                 vector<videoAdsService.getLegalPartnerAdsSettingsRequest>
//    user_id:                  long
//    = videoAdsService.GetLegalPartnerAdsSettingsManyResponse;

//a.inner {I:#} a:I*[int] = a.Inner I;
//a.middle {W:#} {P:Type} {Q:Type} a:(a.inner W) b:(pair P Q) = a.Middle W P Q;
//a.top2 n:# m:# c:vector<a.middle m (pair (a.inner n) (a.inner n)) (a.Inner 3)> = a.Top2;

//usefulService.userEntityPaymentItemPromo#24c7ec9f
//  {fields_mask: #}
//  content: string
//  = usefulService.UserEntityPaymentItemPromo fields_mask;

//usefulService.userEntityPaymentItem#4f798680
//  {fields_mask: #}
//  id: string
////  promo: (Maybe (usefulService.UserEntityPaymentItemPromo fields_mask))
//  = usefulService.UserEntityPaymentItem fields_mask;
//
//usefulService.getUserEntityResult#e494698a
//  {fields_mask: #}
//  entity_id: string
//  payment_items: (Maybe (Dictionary (usefulService.UserEntityPaymentItem fields_mask)))
//  = usefulService.GetUserEntityResult fields_mask;
//
//@readwrite
//usefulService.getUserEntity#3c857e52
//  fields_mask: #
//  stage_id: fields_mask.0?string
//  => Maybe (usefulService.GetUserEntityResult fields_mask);


// comment before Ok
curl.responseOk#57d7a33a
  field_mask:   #
  http_code:    int
  http_headers: %(Dictionary string)
  http_body:    field_mask.0?string
  = curl.Response;

// comment before Error
curl.responseError#af514651
  field_mask:   #
  error_code:   int
  error_string: string
  = curl.Response;

---functions---

// comment before request
@kphp
curl.request#3f5a4651
  field_mask: #
  // comment before
  method:             string
  url:                string // comment right
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
  = curl.Response;---types---
