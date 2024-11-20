cmake_minimum_required(VERSION 3.28)
project(cpp)

set(CMAKE_CXX_STANDARD 20)

add_executable(cpp tests/test_objects_bytes.cpp tests/test_functions_bytes.cpp tests/test_objects_json.cpp)
