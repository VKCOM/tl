boolFalse#bc799737 = Bool;
boolTrue#997275b5 = Bool;
// BASE TYPES
int#a8509bda ? = Int;
long#22076cba ? = Long;

string#b5286e24 ? = String;

//resultFalse#27930a7b {t:Type} = Maybe t;
//resultTrue#3f9c8ef8 {t:Type} t = Maybe t;

pair {X:Type} {Y:Type} a:X b:Y = Pair X Y;

vector#1cb5c415 {t:Type} # [t] = Vector t;
tuple#9770768a {t:Type} {n:#} [t] = Tuple t n;

dictionaryField {t:Type} key:string value:t = DictionaryField t;
dictionary#1f4c618f {t:Type} %(Vector %(DictionaryField t)) = Dictionary t;

//dictionaryAnyField {k:Type} {v:Type} key:k value:v = DictionaryAnyField k v;
//dictionaryAny#1f4c6190 {k:Type} {v:Type} # [(dictionaryAnyField k v)] = DictionaryAny k v;

true = True; // this can be used as void type and serialized to empty array in PHP

map {X:Type} {Y:Type} key:X value:Y = Map X Y;

videoAdsService.getLegalPartnerAdsSettingsManyResponseOk#4db649d8
    fields_mask:              #
    settings:                 vector<videoAdsService.legalPartnerAdsSettings>
    = videoAdsService.GetLegalPartnerAdsSettingsManyResponse;

videoAdsService.getLegalPartnerAdsSettingsManyResponseError#143f3aaa
    fields_mask:              #
    error_code:               int
    error:                    string
    = videoAdsService.GetLegalPartnerAdsSettingsManyResponse;

videoAdsService.legalPartnerAdsSettingsVideoRules#ee6ba4fa
    fields_mask:              #
    video_id:                 int
    //rules:                    videoAdsService.legalPartnerAdsSettingsRules
    = videoAdsService.LegalPartnerAdsSettingsVideoRules;


videoAdsService.legalPartnerAdsSettings#59ee3756
    fields_mask:              #
    //common_rules:             fields_mask.0?videoAdsService.legalPartnerAdsSettingsRules
    video_rules:              fields_mask.1?vector<videoAdsService.legalPartnerAdsSettingsVideoRules>
    //playlist_rules:           fields_mask.2?vector<videoAdsService.legalPartnerAdsSettingsPlaylistRules>
    = videoAdsService.LegalPartnerAdsSettings;


//nonMigrating.example
//    use:               vector<videoAdsService.legalPartnerAdsSettings>
//    use2:              videoAdsService.LegalPartnerAdsSettings
//    = nonMigrating.Example;
//
//nonMigrating2.example
//    use:               videoAdsService.LegalPartnerAdsSettings
//    use2:              vector<videoAdsService.legalPartnerAdsSettings>
//    = nonMigrating2.Example;

---functions---

@read
videoAdsService.getLegalPartnerAdsSettingsMany#98505e85
    fields_mask:              #
//    requests:                 vector<videoAdsService.getLegalPartnerAdsSettingsRequest>
    user_id:                  long
    = videoAdsService.GetLegalPartnerAdsSettingsManyResponse;
