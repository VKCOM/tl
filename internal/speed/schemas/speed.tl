vector#1cb5c415 {t:Type} # [t] = Vector t;

point x:# y:# z:# = Point;
points values:vector<Point> = Points;

memcache.not_found                                                    = memcache.Value;
memcache.strvalue value:string flags:int                              = memcache.Value;
memcache.longvalue value:long flags:int                               = memcache.Value;
memcache.strvalueWithTime value:string flags:int modificationTime:int = memcache.Value;
memcache.longvalueWithTime value:long flags:int modificationTime:int  = memcache.Value;
memcache.strvalueWithDelay value:string flags:int delay:int           = memcache.Value;