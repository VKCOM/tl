CPP_DIR = ../../gen
CASES_DIR = $(CPP_DIR)/cases_cpp
SCHEMA_DIR = $(CPP_DIR)/schema_cpp

CC = g++
CFLAGS = -std=c++20 -O3 -Wno-noexcept-type -g -Wall -Wextra -Werror=return-type -Wno-unused-parameter

all: run-all-tests

# run in project root "make cpp_gen_test_data"
run-all-tests:
	@cd ../../../../../ && make cpp_gen_test_data
	mkdir -p build
	make run-objects-test
	make run-functions-test
	echo "\n\033[0;32mAll tests passed!\033[0m\n"

clean-build:
	rm -rf build
	mkdir build

run-objects-test: build-objects-bytes build-objects-json
	./build/test_objects_bytes
	./build/test_objects_json

run-functions-test: build-functions-bytes
	./build/test_functions_bytes

build-objects-bytes: build/test_objects_bytes.o compile-cases
	$(CC) $(CFLAGS) -g -o build/test_objects_bytes build/test_objects_bytes.o \
		$(CASES_DIR)/build/meta.o \
		$(CASES_DIR)/build/factory.o \
		$(CASES_DIR)/build/basictl.o \
		$(CASES_DIR)/build/string_io.o \
		$(CASES_DIR)/build/tl/__common_namespace.o \
		$(CASES_DIR)/build/tl/benchmarks.o \
		$(CASES_DIR)/build/tl/cases.o \
		$(CASES_DIR)/build/tl/cases_bytes.o

build/test_objects_bytes.o: tests/test_objects_bytes.cpp
	$(CC) $(CFLAGS) -c tests/test_objects_bytes.cpp -o build/test_objects_bytes.o

build-objects-json: build/test_objects_json.o compile-cases
	$(CC) $(CFLAGS) -g -o build/test_objects_json build/test_objects_json.o \
    		$(CASES_DIR)/build/meta.o \
    		$(CASES_DIR)/build/factory.o \
    		$(CASES_DIR)/build/basictl.o \
    		$(CASES_DIR)/build/string_io.o \
    		$(CASES_DIR)/build/tl/__common_namespace.o \
    		$(CASES_DIR)/build/tl/benchmarks.o \
    		$(CASES_DIR)/build/tl/cases.o \
    		$(CASES_DIR)/build/tl/cases_bytes.o

build/test_objects_json.o: tests/test_objects_json.cpp
	$(CC) $(CFLAGS) -c tests/test_objects_json.cpp -o build/test_objects_json.o

compile-cases:
	@cd $(CASES_DIR)/ &&\
 		mkdir -p build &&\
 		make build/basictl.o build/string_io.o build/meta.o build/factory.o build/tl/__common_namespace.o build/tl/benchmarks.o build/tl/cases.o build/tl/cases_bytes.o
	@cd ../../codegen_test/cpp/

build-functions-bytes: build/test_functions_bytes.o compile-schema
	$(CC) $(CFLAGS) -g -o build/test_functions_bytes build/test_functions_bytes.o \
		$(SCHEMA_DIR)/build/meta.o \
		$(SCHEMA_DIR)/build/factory.o \
		$(SCHEMA_DIR)/build/basictl.o \
		$(SCHEMA_DIR)/build/string_io.o \
        $(SCHEMA_DIR)/build/tl/__common_namespace.o \
        $(SCHEMA_DIR)/build/tl/antispam.o \
        $(SCHEMA_DIR)/build/tl/pkg2.o \
        $(SCHEMA_DIR)/build/tl/service1.o \
        $(SCHEMA_DIR)/build/tl/service2.o \
        $(SCHEMA_DIR)/build/tl/service3.o \
        $(SCHEMA_DIR)/build/tl/service4.o \
        $(SCHEMA_DIR)/build/tl/service5.o \
        $(SCHEMA_DIR)/build/tl/service6.o \
        $(SCHEMA_DIR)/build/tl/tasks.o \
        $(SCHEMA_DIR)/build/tl/tree_stats.o \
        $(SCHEMA_DIR)/build/tl/unique.o

build/test_functions_bytes.o: tests/test_functions_bytes.cpp
	$(CC) $(CFLAGS) -c tests/test_functions_bytes.cpp -o build/test_functions_bytes.o

compile-schema:
	@cd $(SCHEMA_DIR)/ &&\
		mkdir -p build &&\
		make build/basictl.o build/string_io.o build/meta.o build/factory.o build/tl/__common_namespace.o build/tl/antispam.o build/tl/pkg2.o build/tl/service1.o build/tl/service2.o build/tl/service3.o build/tl/service4.o build/tl/service5.o build/tl/service6.o build/tl/tasks.o build/tl/tree_stats.o build/tl/unique.o
	@cd ../../codegen_test/cpp/