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


statshouseApi.seriesMeta#5c2bf286 {query_fields_mask:#} fields_mask:#
    time_shift: long
    tags: %(Dictionary string)
    // what: fields_mask.1?statshouseApi.Function
    name: query_fields_mask.4?string
    color: query_fields_mask.5?string
    total: query_fields_mask.6?int
    max_hosts: query_fields_mask.7?%(Vector string)
    = statshouseApi.SeriesMeta query_fields_mask;

statshouseApi.queryResponse#4487e49a {query_fields_mask:#}
    fields_mask:#
    //series: %(statshouseApi.Series)
    series_meta: %(Vector %(statshouseApi.SeriesMeta query_fields_mask))
    chunk_ids: %(Vector int)
    total_time_points: int
    response_id: long
    excess_point_left: fields_mask.0?true
    excess_point_right: fields_mask.1?true
    = statshouseApi.GetQueryResponse query_fields_mask;

---functions---

@kphp
coreEmployees.checkRights#6f40742f
    fields_mask:#
    right_keys: (Vector string)
    user_id: long
    ip_address: fields_mask.0?string
    api_id: fields_mask.1?long
    is_api: fields_mask.2?true
    is_internal_app: fields_mask.3?true
    is_from_cron: fields_mask.4?true
    performer_id: fields_mask.5?long
  = Vector (Map string Bool);


@readwrite
statshouseApi.getQuery#0c7349bb fields_mask:#
    access_token: string
    // query: %statshouseApi.Query
    = statshouseApi.GetQueryResponse fields_mask;
