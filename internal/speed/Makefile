
.PHONY: proto
proto:
	protoc --go_out=../../../../../../.. speed.proto


.PHONY: gen_methods
gen_methods:
	bash scripts/create_proto.sh
	bash scripts/create_tl.sh

COUNT ?= 3

.PHONY: all_benchmarks
all_benchmarks:
	go test -bench=. -count=$(COUNT)

# ----------- POINT ---------------
.PHONY: point_write_gen
point_write_gen:
	go test -bench='^BenchmarkPointWrite.*Gen' -count=$(COUNT)

.PHONY: point_write
point_write:
	go test -bench='^BenchmarkPointWrite' -count=$(COUNT)

.PHONY: point_read_gen
point_read_gen:
	go test -bench='^BenchmarkPointRead.*Gen' -count=$(COUNT)

.PHONY: point_read
point_read:
	go test -bench='^BenchmarkPointRead' -count=$(COUNT)

# ----------- POINTS ---------------
.PHONY: points_read_gen
points_read_gen:
	go test -bench='^BenchmarkPointsRead.*Gen' -count=$(COUNT)

.PHONY: points_write_gen
points_write_gen:
	go test -bench='^BenchmarkPointsWrite.*Gen' -count=$(COUNT)

# ----------- memcache.value ---------------
.PHONY: memcache_value_read_gen
memcache_value_read_gen:
	go test -bench='^BenchmarkMemcacheValueRead.*Gen' -count=$(COUNT)

.PHONY: memcache_value_write_gen
memcache_value_write_gen:
	go test -bench='^BenchmarkMemcacheValueWrite.*Gen' -count=$(COUNT)

# ----------- memcache.values ---------------
.PHONY: memcache_values_read_gen
memcache_values_read_gen:
	go test -bench='^BenchmarkMemcacheValuesRead.*Gen' -count=$(COUNT)

.PHONY: memcache_values_write_gen
memcache_values_write_gen:
	go test -bench='^BenchmarkMemcacheValuesWrite.*Gen' -count=$(COUNT)