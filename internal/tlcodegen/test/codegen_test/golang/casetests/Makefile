run-write-benchmark-array:
	go test \
	-bench=BenchmarkTL.*WriteRandomArrayWithWriteBuffer \
	-count=3 -run=^\# \
	-benchtime=15000000x \
	-benchmem \
	-test.cpuprofile=cpu_array_write.out \
	-test.memprofile=mem_array_write.out \
	-test.outputdir=out

view-write-benchmark-array-cpu:
	go tool pprof -http=localhost:6060 out/cpu_array_write.out

view-write-benchmark-array-mem:
	go tool pprof -http=localhost:6060 out/mem_array_write.out

run-write-benchmark-vector:
	go test \
	-bench=BenchmarkTL.*WriteRandomVectorWithWriteBuffer \
	-count=2 -run=^\# \
	-benchtime=300000x \
	-benchmem \
	-test.cpuprofile=cpu_vector_write.out \
	-test.memprofile=mem_vector_write.out \
	-test.outputdir=out

view-write-benchmark-vector-cpu:
	go tool pprof -http=localhost:6060 out/cpu_vector_write.out

view-write-benchmark-vector-mem:
	go tool pprof -http=localhost:6060 out/mem_vector_write.out

run-read-benchmark-vector:
	go test \
	-bench=BenchmarkTL.*ReadRandomVector \
	-count=3 -run=^\# \
	-benchtime=45000000x \
	-benchmem \
	-test.cpuprofile=cpu_vector_read.out \
	-test.memprofile=mem_vector_read.out \
	-test.outputdir=out

view-read-benchmark-vector-cpu:
	go tool pprof -http=localhost:6060 out/cpu_vector_read.out

view-read-benchmark-vector-mem:
	go tool pprof -http=localhost:6060 out/mem_vector_read.out