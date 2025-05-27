CPP_DIR = ../../gen
CASES_DIR = $(CPP_DIR)/cases_cpp
SCHEMA_DIR = $(CPP_DIR)/schema_cpp
GOLDMASTER_DIR = $(CPP_DIR)/goldmaster_cpp

CC = g++
CFLAGS = -std=c++20 -O3 -Wno-noexcept-type -g -Wall -Wextra -Werror=return-type -Wno-unused-parameter

build-test-gen:
	@cd ../../../../../; \
	$(MAKE) build; \
	$(MAKE) cpp_gen_test_data;

# run in project root "make cpp_gen_test_data"
run-all-tests: build-test-gen
	make clean-build; \
	make run-objects-test; \
	make run-functions-test; \
	make run-goldmaster-stress;

clean-build:
	rm -rf build
	mkdir build

run-objects-test: build-objects-bytes
	./build/test_objects_bytes

run-functions-test: build-functions-bytes
	./build/test_functions_bytes

run-goldmaster-stress: build-goldmaster-stress
	./build/test_goldmaster_stress

build-objects-bytes: build/test_objects_bytes.o compile-cases
	$(CC) $(CFLAGS) -o build/test_objects_bytes build/test_objects_bytes.o \
		$(CASES_DIR)/__build/io_throwable_streams.o \
		$(CASES_DIR)/__build/io_streams.o \
		$(CASES_DIR)/__build/string_io.o \
		$(CASES_DIR)/__build/__meta.o \
		$(CASES_DIR)/__build/__factory.o \
		$(CASES_DIR)/__build/__common_namespace.o \
		$(CASES_DIR)/__build/benchmarks.o \
		$(CASES_DIR)/__build/cases.o \
		$(CASES_DIR)/__build/cases_bytes.o \
		$(CASES_DIR)/__build/casesTL2.o

build/test_objects_bytes.o: tests/test_objects_bytes.cpp
	$(CC) $(CFLAGS) -I$(CASES_DIR) -c tests/test_objects_bytes.cpp -o build/test_objects_bytes.o

compile-cases:
	@cd $(CASES_DIR)/ && make __build/io_streams.o __build/io_throwable_streams.o __build/string_io.o __build/__common_namespace.o __build/benchmarks.o __build/cases.o __build/cases_bytes.o __build/casesTL2.o __build/__meta.o __build/__factory.o
	@cd ../../codegen_test/cpp/

build-functions-bytes: build/test_functions_bytes.o compile-schema
	$(CC) $(CFLAGS) -o build/test_functions_bytes build/test_functions_bytes.o \
		$(SCHEMA_DIR)/__build/io_throwable_streams.o \
		$(SCHEMA_DIR)/__build/io_streams.o \
		$(SCHEMA_DIR)/__build/string_io.o \
		$(SCHEMA_DIR)/__build/__meta.o \
		$(SCHEMA_DIR)/__build/__factory.o \
        $(SCHEMA_DIR)/__build/__common_namespace.o \
        $(SCHEMA_DIR)/__build/antispam.o \
        $(SCHEMA_DIR)/__build/pkg2.o \
        $(SCHEMA_DIR)/__build/service1.o \
        $(SCHEMA_DIR)/__build/service2.o \
        $(SCHEMA_DIR)/__build/service3.o \
        $(SCHEMA_DIR)/__build/service4.o \
        $(SCHEMA_DIR)/__build/service5.o \
        $(SCHEMA_DIR)/__build/service6.o \
        $(SCHEMA_DIR)/__build/tasks.o \
        $(SCHEMA_DIR)/__build/tree_stats.o \
        $(SCHEMA_DIR)/__build/unique.o

build/test_functions_bytes.o: tests/test_functions_bytes.cpp compile-schema
	$(CC) $(CFLAGS) -I$(SCHEMA_DIR) -c tests/test_functions_bytes.cpp -o build/test_functions_bytes.o

compile-schema:
	@cd $(SCHEMA_DIR)/ && make __build/io_streams.o __build/io_throwable_streams.o __build/string_io.o __build/__common_namespace.o __build/antispam.o __build/pkg2.o __build/service1.o __build/service2.o __build/service3.o __build/service4.o __build/service5.o __build/service6.o __build/tasks.o __build/tree_stats.o __build/unique.o __build/__meta.o __build/__factory.o
	@cd ../../codegen_test/cpp/


build-goldmaster-stress: build/test_goldmaster_stress.o compile-goldmaster
	$(CC) $(CFLAGS) -o build/test_goldmaster_stress build/test_goldmaster_stress.o \
		$(GOLDMASTER_DIR)/__build/io_streams.o \
		$(GOLDMASTER_DIR)/__build/io_throwable_streams.o \
		$(GOLDMASTER_DIR)/__build/string_io.o \
		$(GOLDMASTER_DIR)/__build/__meta.o \
		$(GOLDMASTER_DIR)/__build/__factory.o \
		$(GOLDMASTER_DIR)/__build/__common_namespace.o \
		$(GOLDMASTER_DIR)/__build/a.o \
		$(GOLDMASTER_DIR)/__build/ab.o \
		$(GOLDMASTER_DIR)/__build/au.o \
		$(GOLDMASTER_DIR)/__build/cd.o \
		$(GOLDMASTER_DIR)/__build/cyc1.o \
		$(GOLDMASTER_DIR)/__build/cyc2.o \
		$(GOLDMASTER_DIR)/__build/cyc3.o \
		$(GOLDMASTER_DIR)/__build/service5.o \
		$(GOLDMASTER_DIR)/__build/service5Long.o \
		$(GOLDMASTER_DIR)/__build/usefulService.o

build/test_goldmaster_stress.o: tests/stress_test_goldmaster.cpp compile-goldmaster
	$(CC) $(CFLAGS) -I$(GOLDMASTER_DIR) -c tests/stress_test_goldmaster.cpp -o build/test_goldmaster_stress.o

compile-goldmaster:
	@cd $(GOLDMASTER_DIR)/ && make __build/io_streams.o __build/io_throwable_streams.o __build/string_io.o __build/__common_namespace.o __build/a.o __build/ab.o __build/au.o __build/cd.o __build/cyc1.o __build/cyc2.o __build/cyc3.o __build/service5.o __build/service5Long.o __build/usefulService.o  __build/__meta.o __build/__factory.o
	@cd ../../codegen_test/cpp/