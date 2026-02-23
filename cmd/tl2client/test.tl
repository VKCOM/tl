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

---types---

reqError#b527877d {X:Type} error_code:int error:string = ReqResult X;
reqResultHeader#8cc84ce1 {X:Type} flags:# extra:%(RpcReqResultExtra flags) result:X = ReqResult X;
//reqResultFalse {X:Type} = ReqResult X;
_ {X:Type} result:X = ReqResult X;

rpcReqResult#63aeda4e {X:Type} query_id:long result:(ReqResult X) = RpcReqResult X;
rpcReqError#7ae432f5 {X:Type} query_id:long error_code:int error:string = RpcReqResult X;

rpcPong#8430eaa7 ping_id:long = RpcPong;

left {X:Type} {Y:Type} value:X = Either X Y;
right {X:Type} {Y:Type} value:Y = Either X Y;

concat {X:Type} {Y:Type} x:X y:Y = Concat X Y;

// Experimental. Top-level packet magic. Client sends on request cancel, so server can free context associated with request.
// No response from server is sent. Required for correctly cancelling long poll requests.
// Supported by vkgo rpc.Server/rpc.Client for now. Support is advertised during PacketConn handshake with flag bit.
rpcCancelReq#193f1b22 query_id:long = RpcCancelReq;

// Top-level packet magic. Server sends RpcServerWantsFin once per connection to ask for gracefully closing connection.
// Client should finish sending current request, send RpcClientWantsFin, send no more requests and wait connection to be closed by server.
// Server will send all non-longpoll responses, then close connection. Client will send all new requests to a new connection.
// Long poll requests make this protocol more complicated, as client does not know which requests have been added to long poll map,
// hence client cannot rely on waiting to finish all responses.
// Supported by vkgo rpc.Server/rpc.Client for now. Support is advertised during PacketConn handshake with flag bit.
rpcServerWantsFin#a8ddbc46 = RpcServerWantsFin;
rpcClientWantsFin#0b73429e = RpcClientWantsFin;

// We wrap body of TL2 request with this marker.
// Any response is parsed in the context of request, so we decided to not wrap response.
// We will remove all wrapping later as part of rpcInvokeReq2 transition.
// We do not want to define this marker as wrapper, because we never want it to be parsed by vkext and factories.
// TL1 basically breaks after this marker.
rpcTL2Marker#30324c54 = RpcTL2Marker;

---functions---

@any @internal rpcPing#5730a2df ping_id:long = RpcPong;

@any rpcDestActor#7568aabd {X:Type} actor_id:long query:!X = X;
@any rpcDestActorFlags#f0a5acf7 {X:Type} actor_id:long flags:# extra:%(RpcInvokeReqExtra flags) query:!X = X;
@any rpcDestFlags#e352035e {X:Type} flags:# extra:%(RpcInvokeReqExtra flags) query:!X = X;

@any rpcInvokeReq#2374df3d {X:Type} query_id:long query:!X = RpcReqResult X;

// The only allowed grammar of invoking request is:
// rpcInvokeReq (rpcDestActor | rpcDestActorFlags | rpcDestFlags)? rpcTL2Marker? Body


---types---

engine.query {X:Type} query:!X = engine.Query;
engine.queryShortened query:%(VectorTotal int) = engine.Query;


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

videoAdsService.adsParams#173aa2ee
    fields_mask:              #
    // post:                     fields_mask.0?videoAdsService.fullPostID
    is_embed:                 fields_mask.1?true
    promo_post_hash:          fields_mask.2?string
    trach_code:               fields_mask.3?string
    device_id:                fields_mask.4?string
    query_params:             fields_mask.5?dictionary<string>
    = videoAdsService.AdsParams;

---functions---

@read
videoAdsService.checkMany#daf58753
    fields_mask:              #
    params:                   vector<videoAdsService.checkParams>
    device_id:                fields_mask.0?string
    platform:                 fields_mask.1?string
    api_version:              fields_mask.2?int
//    user_info:                fields_mask.3?%videoUserInfo.info
    = Vector videoAdsService.CheckResponse;

@readwrite
videoApi.getByIds#cf03372f
    fields_mask:  #
//    model_params: fields_mask.0?%videoApi.GetByIdsParams
    ads:          fields_mask.1?%videoAdsService.AdsParams
    = Int;
