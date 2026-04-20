
.PHONY: proto
proto:
	protoc --go_out=../../../../../../.. speed.proto

.PHONY: point_write_gen
point_write_gen:
	go test -bench='^BenchmarkPointWrite.*Gen'

.PHONY: point_read_gen
point_read_gen:
	go test -bench='^BenchmarkPointRead.*Gen'

.PHONY: point_read
point_read:
	go test -bench='^BenchmarkPointRead'
