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

//dictionaryField {t:Type} key:string value:t = DictionaryField t;
//dictionary#1f4c618f {t:Type} %(Vector %(DictionaryField t)) = Dictionary t;

//dictionaryAnyField {k:Type} {v:Type} key:k value:v = DictionaryAnyField k v;
//dictionaryAny#1f4c6190 {k:Type} {v:Type} # [(dictionaryAnyField k v)] = DictionaryAny k v;

true = True; // this can be used as void type and serialized to empty array in PHP

---types---


videoCommon.fullVideoID#b45b85e1
    fields_mask:    #
    video_id:       fields_mask.0?int
    owner_id:       fields_mask.1?long
= videoCommon.FullVideoID;


videoAdsService.checkParams#6742a1bb
    fields_mask:              #
//    fields:                   fields_mask.0?videoAdsService.checkFields
    full_video_id:            fields_mask.1?videoCommon.fullVideoID
    user_id:                  fields_mask.2?long
    is_child_account:         fields_mask.3?true
    = videoAdsService.CheckParams;

videoAdsService.usedType x:int = videoAdsService.UsedType;

videoAdsService.user
    x:pair<vector<videoAdsService.usedType>, int>
    y:Pair<Vector<videoAdsService.usedType>, int>
   = videoAdsService.User;

videoAdsService.snackbarEventTypeClosed#d5f13be8     = videoAdsService.SnackbarEventType;
videoAdsService.snackbarEventTypeSubscribed#052dccb0 = videoAdsService.SnackbarEventType;

//videoUserInfo.info#78e87a2f
//    fields_mask:            #
//    user_id:                long
//    active_subscription:    fields_mask.0?true
//    child_mode_active:      fields_mask.1?true
//    is_child_account:       fields_mask.2?true
//    is_vk_edu:              fields_mask.3?true
//= videoUserInfo.Info;
//
//videoAdsService.checkParams#6742a1bb
//    fields_mask:              #
//    fields:                   fields_mask.0?videoAdsService.checkFields
//    full_video_id:            fields_mask.1?videoCommon.fullVideoID
//    user_id:                  fields_mask.2?long
//    = videoAdsService.CheckParams;
//
//
//videoAdsService.checkFields#cf0cd638
//    fields_mask:                  #
//    can_show_banner:              fields_mask.0?true
//    ads_disabled:                 fields_mask.1?true
//    = videoAdsService.CheckFields;

videoAdsService.checkResponseOk#87fad1f4
    fields_mask:              #
    can_show_banner:                  fields_mask.0?true
    ads_disabled:                     fields_mask.1?true
    = videoAdsService.CheckResponse;

videoAdsService.checkResponseError#83eb9781
    fields_mask:              #
    error_code:               int
    error:                    string
    = videoAdsService.CheckResponse;

//---functions---
//
//@read
//videoAdsService.checkMany#daf58753
//    fields_mask:              #
//    params:                   vector<videoAdsService.checkParams>
//    device_id:                fields_mask.0?string
//    platform:                 fields_mask.1?string
//    api_version:              fields_mask.2?int
//    user_info:                fields_mask.3?%videoUserInfo.info
//    = Vector videoAdsService.CheckResponse;


benchmarks.vruposition#32792c04 fields_mask:#
    commit_bit:         fields_mask.0?true // we set this bit only if block with this bit set is committed, so we have no more than 1 such uncommitted block.
    file_offset:        long
    seq_number:         fields_mask.14?long // incremented after each block.
= benchmarks.VruPosition;

benchmarks.vrutoytopLevelUnionBig
    next_positions:(vector benchmarks.vruposition)
= benchmarks.VrutoyTopLevelUnion;

benchmarks.vrutoytopLevelUnionEmpty
= benchmarks.VrutoyTopLevelUnion;
