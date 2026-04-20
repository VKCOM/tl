
.PHONY: proto
proto:
	protoc --go_out=../../../../../../.. speed.proto

.PHONY: point_write_gen
point_write_gen:
	go test -bench='^BenchmarkPointWrite.*Gen'
