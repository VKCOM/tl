
.PHONY: proto
proto:
	protoc --go_out=../../../../../../.. speed.proto

COUNT ?= 3

.PHONY: all_benchmarks
all_benchmarks:
	go test -bench=. -count=$(COUNT)

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
