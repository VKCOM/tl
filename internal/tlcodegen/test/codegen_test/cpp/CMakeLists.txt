cmake_minimum_required(VERSION 3.28)
project(cpp_tl_test)

set(CMAKE_CXX_STANDARD 20)

add_executable(test_objects_json
        tests/test_objects_json.cpp
#        tests/test_objects_bytes.cpp
#        tests/test_functions_bytes.cpp
		../../gen/cases_cpp/basictl/impl/string_io.h
		../../gen/cases_cpp/basictl/impl/string_io.cpp
		../../gen/cases_cpp/basictl/io_connectors.h
		../../gen/cases_cpp/basictl/errors.h
		../../gen/cases_cpp/basictl/constants.h
		../../gen/cases_cpp/basictl/io_streams.cpp
		../../gen/cases_cpp/basictl/io_throwable_streams.cpp

		../../gen/cases_cpp/__meta/details.cpp
		../../gen/cases_cpp/__factory/details.cpp

		../../gen/cases_cpp/__common_namespace/details.cpp
		../../gen/cases_cpp/benchmarks/details.cpp
		../../gen/cases_cpp/cases/details.cpp
		../../gen/cases_cpp/cases_bytes/details.cpp
)

add_executable(test_objects_bytes
#        tests/test_objects_json.cpp
        tests/test_objects_bytes.cpp
#        tests/test_functions_bytes.cpp
        ../../gen/cases_cpp/basictl/impl/string_io.h
        ../../gen/cases_cpp/basictl/impl/string_io.cpp
        ../../gen/cases_cpp/basictl/io_connectors.h
        ../../gen/cases_cpp/basictl/errors.h
        ../../gen/cases_cpp/basictl/constants.h
		../../gen/cases_cpp/basictl/io_streams.cpp
		../../gen/cases_cpp/basictl/io_throwable_streams.cpp

		../../gen/cases_cpp/__meta/details.cpp
        ../../gen/cases_cpp/__factory/details.cpp

		../../gen/cases_cpp/__common_namespace/details.cpp
		../../gen/cases_cpp/benchmarks/details.cpp
		../../gen/cases_cpp/cases/details.cpp
		../../gen/cases_cpp/cases_bytes/details.cpp
)

add_executable(test_functions_bytes
#        tests/test_objects_json.cpp
#        tests/test_objects_bytes.cpp
        tests/test_functions_bytes.cpp
        ../../gen/schema_cpp/basictl/impl/string_io.h
        ../../gen/schema_cpp/basictl/impl/string_io.cpp
        ../../gen/schema_cpp/basictl/io_connectors.h
        ../../gen/schema_cpp/basictl/errors.h
        ../../gen/schema_cpp/basictl/constants.h
		../../gen/schema_cpp/basictl/io_streams.cpp
		../../gen/schema_cpp/basictl/io_throwable_streams.cpp

		../../gen/schema_cpp/__meta/details.cpp
		../../gen/schema_cpp/__factory/details.cpp

        ../../gen/schema_cpp/__common_namespace/details.cpp
        ../../gen/schema_cpp/antispam/details.cpp
        ../../gen/schema_cpp/pkg2/details.cpp
        ../../gen/schema_cpp/service1/details.cpp
        ../../gen/schema_cpp/service2/details.cpp
        ../../gen/schema_cpp/service3/details.cpp
        ../../gen/schema_cpp/service4/details.cpp
        ../../gen/schema_cpp/service5/details.cpp
        ../../gen/schema_cpp/service6/details.cpp
        ../../gen/schema_cpp/tasks/details.cpp
        ../../gen/schema_cpp/tree_stats/details.cpp
        ../../gen/schema_cpp/unique/details.cpp
)


add_executable(stress_test_goldmaster
		#        tests/test_objects_json.cpp
		#        tests/test_objects_bytes.cpp
		tests/stress_test_goldmaster.cpp
		../../gen/goldmaster_cpp/basictl/impl/string_io.h
		../../gen/goldmaster_cpp/basictl/impl/string_io.cpp
		../../gen/goldmaster_cpp/basictl/io_connectors.h
		../../gen/goldmaster_cpp/basictl/errors.h
		../../gen/goldmaster_cpp/basictl/constants.h
		../../gen/goldmaster_cpp/basictl/io_streams.cpp
		../../gen/goldmaster_cpp/basictl/io_throwable_streams.cpp

		../../gen/goldmaster_cpp/__meta/details.cpp
		../../gen/goldmaster_cpp/__factory/details.cpp

		../../gen/goldmaster_cpp/__common_namespace/details.cpp
		../../gen/goldmaster_cpp/a/details.cpp
		../../gen/goldmaster_cpp/ab/details.cpp
		../../gen/goldmaster_cpp/au/details.cpp
		../../gen/goldmaster_cpp/cd/details.cpp
		../../gen/goldmaster_cpp/cyc1/details.cpp
		../../gen/goldmaster_cpp/cyc2/details.cpp
		../../gen/goldmaster_cpp/cyc3/details.cpp
		../../gen/goldmaster_cpp/service5/details.cpp
		../../gen/goldmaster_cpp/service5Long/details.cpp
		../../gen/goldmaster_cpp/usefulService/details.cpp
)