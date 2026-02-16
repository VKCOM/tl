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

dictionaryAnyField {k:Type} {v:Type} key:k value:v = DictionaryAnyField k v;
dictionaryAny#1f4c6190 {k:Type} {v:Type} # [(dictionaryAnyField k v)] = DictionaryAny k v;

true = True; // this can be used as void type and serialized to empty array in PHP

---types---


videoCommon.fullVideoID#b45b85e1
    fields_mask:    #
    video_id:       fields_mask.0?int
    owner_id:       fields_mask.1?long
= videoCommon.FullVideoID;

videoUserInfo.info#78e87a2f
    fields_mask:            #
    user_id:                long
    active_subscription:    fields_mask.0?true
    child_mode_active:      fields_mask.1?true
    is_child_account:       fields_mask.2?true
    is_vk_edu:              fields_mask.3?true
= videoUserInfo.Info;

videoAdsService.snackbarEventTypeClosed#d5f13be8     = videoAdsService.SnackbarEventType;
videoAdsService.snackbarEventTypeSubscribed#052dccb0 = videoAdsService.SnackbarEventType;

videoAdsService.fullPostID#4dda7b2f
    fields_mask:              #
    owner_id:                 long
    post_id:                  long
    = videoAdsService.FullPostID;

videoAdsService.checkParams#6742a1bb
    fields_mask:              #
    fields:                   fields_mask.0?videoAdsService.checkFields
    full_video_id:            fields_mask.1?videoCommon.fullVideoID
    user_id:                  fields_mask.2?long
    = videoAdsService.CheckParams;

videoAdsService.checkFields#cf0cd638
    fields_mask:                  #
    can_show_banner:              fields_mask.0?true
    ads_disabled:                 fields_mask.1?true
    legal_partner:                fields_mask.2?videoAdsService.legalPartnerCheckFields
    ads_type:                     fields_mask.3?videoAdsService.adsTypeCheckFields
    = videoAdsService.CheckFields;

videoAdsService.checkResponseOk#87fad1f4
    fields_mask:              #
    can_show_banner:                  fields_mask.0?true
    ads_disabled:                     fields_mask.1?true
    legal_partner:                    fields_mask.2?videoAdsService.legalPartnerCheckResult
    ads_type:                         fields_mask.3?videoAdsService.adsTypeCheckResult
    = videoAdsService.CheckResponse;

videoAdsService.checkResponseError#83eb9781
    fields_mask:              #
    error_code:               int
    error:                    string
    = videoAdsService.CheckResponse;

videoAdsService.legalPartnerCheckFields#4b1ba82e
    fields_mask:                  #
    is_legal_partner_video:       fields_mask.0?true
    cat_id:                       fields_mask.1?true
    live_ads_enabled:             fields_mask.2?true
    = videoAdsService.LegalPartnerCheckFields;

videoAdsService.adsTypeCheckFields#298ccbcd
    fields_mask:                  #
    cat_id:                       fields_mask.0?true
    = videoAdsService.AdsTypeCheckFields;

videoAdsService.legalPartnerCheckResult#e66acc5f
    fields_mask:                  #
    is_legal_partner_video:       fields_mask.0?true
    cat_id:                       fields_mask.1?int
    live_ads_enabled:             fields_mask.2?true
    = videoAdsService.LegalPartnerCheckResult;

videoAdsService.adsTypeCheckResult#e946e17a
    fields_mask:                  #
    cat_id:                       fields_mask.0?int
    = videoAdsService.AdsTypeCheckResult;

videoAdsService.ad#6f8e4ffb
    fields_mask:              #
    type:                     fields_mask.0?string
    is_midroll_point:         fields_mask.1?true
    do_after:                 fields_mask.2?true
    point:                    fields_mask.3?int
    url:                      fields_mask.4?string
    = videoAdsService.Ad;

videoAdsService.section#a3f8c2d4
    fields_mask:              #
    key:                      fields_mask.0?string
    ads:                      fields_mask.1?vector<videoAdsService.ad>
    = videoAdsService.Section;

videoAdsService.instream#e7b9a5f1
    fields_mask:              #
    sections:                 fields_mask.0?vector<videoAdsService.section>
    = videoAdsService.Instream;

videoAdsService.sport#9d4e6b8c
    fields_mask:              #
    sections:                 fields_mask.0?vector<videoAdsService.section>
    = videoAdsService.Sport;

videoAdsService.mobile#e9428ac0
    fields_mask:              #
    version:                  fields_mask.0?string
    sections:                 fields_mask.1?vector<videoAdsService.section>
    = videoAdsService.Mobile;

videoAdsService.overlay#1eda475f
    fields_mask:              #
    start_at:                 fields_mask.0?int
    banner_id:                fields_mask.1?string
    = videoAdsService.Overlay;

videoAdsService.pause#1eda4760
    fields_mask:              #
    show_delay:               fields_mask.0?int
    banner_id:                fields_mask.1?string
    interval:                 fields_mask.2?videoAdsService.pauseInterval
    = videoAdsService.Pause;

videoAdsService.pauseInterval#1eda4761
    fields_mask:              #
    start:                    int
    end:                      int
    = videoAdsService.PauseInterval;

videoAdsService.legalPartnerAdsSettings#59ee3756
    fields_mask:              #
    common_rules:             fields_mask.0?videoAdsService.legalPartnerAdsSettingsRules
    video_rules:              fields_mask.1?vector<videoAdsService.legalPartnerAdsSettingsVideoRules>
    playlist_rules:           fields_mask.2?vector<videoAdsService.legalPartnerAdsSettingsPlaylistRules>
    = videoAdsService.LegalPartnerAdsSettings;

videoAdsService.legalPartnerAdsSettingsVideoRules#ee6ba4fa
    fields_mask:              #
    video_id:                 int
    rules:                    videoAdsService.legalPartnerAdsSettingsRules
    = videoAdsService.LegalPartnerAdsSettingsVideoRules;

videoAdsService.legalPartnerAdsSettingsPlaylistRules#72c0b299
    fields_mask:              #
    playlist_id:              int
    rules:                    videoAdsService.legalPartnerAdsSettingsRules
    = videoAdsService.LegalPartnerAdsSettingsPlaylistRules;

videoAdsService.legalPartnerAdsSettingsRules#12a8c9cf
    fields_mask:              #
    cat_id:                   fields_mask.0?int
    live_ads_enabled:         fields_mask.1?Bool // tlgen:nolint
    = videoAdsService.LegalPartnerAdsSettingsRules;

videoAdsService.getAdsResponseOk#3309dce6
    fields_mask:              #
    instream:                 fields_mask.0?videoAdsService.instream
    sport:                    fields_mask.1?videoAdsService.sport
    mobile:                   fields_mask.2?videoAdsService.mobile
    overlays:                 fields_mask.3?vector<videoAdsService.overlay>
    pauses:                   fields_mask.4?vector<videoAdsService.pause>
    = videoAdsService.GetAdsResponse;

videoAdsService.getAdsResponseError#9ea73da5
    error_code:               int
    error:                    string
    = videoAdsService.GetAdsResponse;

videoAdsService.trackAdEventResponseOk#f0523127
    fields_mask:              #
    = videoAdsService.TrackAdEventResponse;

videoAdsService.trackAdEventResponseError#7829da87
    fields_mask:              #
    error_code:               int
    error:                    string
    = videoAdsService.TrackAdEventResponse;

videoAdsService.getSnackbarInfoResponseOk#740086e7
    fields_mask:              #
    can_show:                 fields_mask.0?true
    text:                     fields_mask.1?string
    url:                      fields_mask.2?string
    ttl_seconds:              fields_mask.3?int
    delay_seconds:            fields_mask.4?int
    = videoAdsService.GetSnackbarInfoResponse;

videoAdsService.getSnackbarInfoResponseError#d3102eee
    fields_mask:              #
    error_code:               int
    error:                    string
    = videoAdsService.GetSnackbarInfoResponse;

videoAdsService.trackSnackbarEventResponseOk#ffc7f550
    fields_mask:              #
    = videoAdsService.TrackSnackbarEventResponse;

videoAdsService.trackSnackbarEventResponseError#f6e5e81e
    fields_mask:              #
    error_code:               int
    error:                    string
    = videoAdsService.TrackSnackbarEventResponse;

videoAdsService.updateUserResponseOk#21d76976
    fields_mask:              #
    = videoAdsService.UpdateUserResponse;

videoAdsService.updateUserResponseError#95d05bd8
    fields_mask:              #
    error_code:               int
    error:                    string
    = videoAdsService.UpdateUserResponse;

videoAdsService.adsParams#173aa2ee
    fields_mask:              #
    post:                     fields_mask.0?videoAdsService.fullPostID
    is_embed:                 fields_mask.1?true
    promo_post_hash:          fields_mask.2?string
    trach_code:               fields_mask.3?string
    device_id:                fields_mask.4?string
    query_params:             fields_mask.5?dictionary<string>
    = videoAdsService.AdsParams;

videoAdsService.adsDataOverlay#f9335612
    fields_mask:                #
    slot_id:                    fields_mask.0?int
    duration_seconds:           fields_mask.1?int
    close_button_delay_seconds: fields_mask.2?int
    = videoAdsService.AdsDataOverlay;

videoAdsService.adsDataPause#f9335613
    fields_mask:                #
    slot_id:                    fields_mask.0?int
    duration_seconds:           fields_mask.1?int
    close_button_delay_seconds: fields_mask.2?int
    next_show_after_seconds:    fields_mask.3?int
    = videoAdsService.AdsDataPause;

videoAdsService.getAdsDataResponseOk#3ef5f1e1
    fields_mask:              #
    overlays:                 fields_mask.0?vector<videoAdsService.adsDataOverlay>
    pauses:                   fields_mask.1?vector<videoAdsService.adsDataPause>
    = videoAdsService.GetAdsDataResponse;

videoAdsService.getAdsDataResponseError#bac0e099
    fields_mask:              #
    error_code:               int
    error:                    string
    = videoAdsService.GetAdsDataResponse;

videoAdsService.getLegalPartnerAdsSettingsRequest#4ec07ced
    fields_mask:              #
    owner_id:                 long
    = videoAdsService.GetLegalPartnerAdsSettingsRequest;

videoAdsService.getLegalPartnerAdsSettingsManyResponseOk#4db649d8
    fields_mask:              #
    settings:                 vector<videoAdsService.legalPartnerAdsSettings>
    = videoAdsService.GetLegalPartnerAdsSettingsManyResponse;

videoAdsService.getLegalPartnerAdsSettingsManyResponseError#143f3aaa
    fields_mask:              #
    error_code:               int
    error:                    string
    = videoAdsService.GetLegalPartnerAdsSettingsManyResponse;

videoAdsService.setLegalPartnerAdsSettingsResponseOk#02ab245c
    fields_mask:              #
    = videoAdsService.SetLegalPartnerAdsSettingsResponse;

videoAdsService.setLegalPartnerAdsSettingsResponseError#705cbb70
    fields_mask:              #
    error_code:               int
    error:                    string
    = videoAdsService.SetLegalPartnerAdsSettingsResponse;

videoAdsService.deleteLegalPartnerAdsSettingsResponseOk#5294284b
    fields_mask:              #
    = videoAdsService.DeleteLegalPartnerAdsSettingsResponse;

videoAdsService.deleteLegalPartnerAdsSettingsResponseError#ff421268
    fields_mask:              #
    error_code:               int
    error:                    string
    = videoAdsService.DeleteLegalPartnerAdsSettingsResponse;

---functions---

@read
videoAdsService.getAds#c6596b41
    fields_mask:              #
    full_video_id:            fields_mask.0?videoCommon.fullVideoID
    app_id:                   fields_mask.1?int
    user_id:                  fields_mask.2?long
    full_post_id:             fields_mask.3?videoAdsService.fullPostID
    platform:                 fields_mask.4?string
    is_embed:                 fields_mask.5?true
    api_version:              fields_mask.6?int
    build:                    fields_mask.7?int
    lang_id:                  fields_mask.8?int
    user_agent:               fields_mask.9?string
    device_id:                fields_mask.10?string
    track_code:               fields_mask.11?string
    user_ip:                  fields_mask.12?long
    promo_post_hash:          fields_mask.13?string
    query_params:             fields_mask.14?dictionary<string>
    anonym_id:                fields_mask.15?long
    user_info:                fields_mask.16?%videoUserInfo.info
    = videoAdsService.GetAdsResponse;

@write
videoAdsService.trackAdEvent#b8ca1a4c
    fields_mask:              #
    banner_id:                fields_mask.0?string
    full_video_id:            fields_mask.1?videoCommon.fullVideoID
    user_id:                  fields_mask.2?long
    anonym_id:                fields_mask.3?long
    device_id:                fields_mask.4?string
    section_key:              fields_mask.5?string
    = videoAdsService.TrackAdEventResponse;

@read
videoAdsService.checkMany#daf58753
    fields_mask:              #
    params:                   vector<videoAdsService.checkParams>
    device_id:                fields_mask.0?string
    platform:                 fields_mask.1?string
    api_version:              fields_mask.2?int
    user_info:                fields_mask.3?%videoUserInfo.info
    = Vector videoAdsService.CheckResponse;

@read
videoAdsService.getSnackbarInfo#21d54d5b
    fields_mask:              #
    user_id:                  long
    platform:                 fields_mask.0?string
    user_info:                fields_mask.1?%videoUserInfo.info
    = videoAdsService.GetSnackbarInfoResponse;

@write
videoAdsService.trackSnackbarEvent#d647c937
    fields_mask:              #
    user_id:                  long
    event:                    fields_mask.0?videoAdsService.SnackbarEventType
    = videoAdsService.TrackSnackbarEventResponse;

@read
videoAdsService.getAdsData#773d79c2
    fields_mask:              #
    banner_ids:               fields_mask.0?(vector string)
    user_id:                  fields_mask.1?long
    anonym_id:                fields_mask.2?long
    device_id:                fields_mask.3?string
    platform:                 fields_mask.4?string
    = videoAdsService.GetAdsDataResponse;

@write
videoAdsService.updateUser#2a34a2e4
    fields_mask:              #
    user_id:                  long
    ads_disabled:             fields_mask.0?Bool // tlgen:nolint
    = videoAdsService.UpdateUserResponse;

@read
videoAdsService.getLegalPartnerAdsSettingsMany#98505e85
    fields_mask:              #
    requests:                 vector<videoAdsService.getLegalPartnerAdsSettingsRequest>
    user_id:                  long
    = videoAdsService.GetLegalPartnerAdsSettingsManyResponse;

@write
videoAdsService.setLegalPartnerAdsSettings#9d7f698c
    fields_mask:              #
    owner_id:                 long
    settings:                 videoAdsService.legalPartnerAdsSettings
    user_id:                  long
    = videoAdsService.SetLegalPartnerAdsSettingsResponse;

@write
videoAdsService.deleteLegalPartnerAdsSettings#2592a274
    fields_mask:              #
    owner_id:                 long
    user_id:                  long
    = Int;---types---
