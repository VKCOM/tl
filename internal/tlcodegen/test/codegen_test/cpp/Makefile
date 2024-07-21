CPP_DIR = ../../gen/cases_cpp

CC = g++
CFLAGS = -std=c++17 -O3 -Wno-noexcept-type -g -Wall -Wextra -Werror=return-type -Wno-unused-parameter

run-test: compile-test
	./build/test-bytes

compile-test: test-bytes.o
	$(CC) $(CFLAGS) -o build/test-bytes build/test-bytes.o \
		$(CPP_DIR)/__build/__common.o \
		$(CPP_DIR)/__build/__independent_types.o \
		$(CPP_DIR)/__build/benchmarks.o \
		$(CPP_DIR)/__build/cases.o \
		$(CPP_DIR)/__build/cases_bytes.o

build-schema:
	@cd $(CPP_DIR)/ && make __build/__common.o __build/__independent_types.o __build/benchmarks.o __build/cases.o __build/cases_bytes.o
	@cd ../../codegen_test/cpp/

test-bytes.o: test-bytes.cpp
	$(CC) $(CFLAGS) -c test-bytes.cpp -o build/test-bytes.o