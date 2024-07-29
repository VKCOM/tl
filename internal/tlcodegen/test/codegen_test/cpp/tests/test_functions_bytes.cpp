#include <iostream>
#include <fstream>
#include <bitset>

#include "../utils/hex.h"
#include "../dependencies/json.hpp"

#include "../../../gen/schema_cpp/a_tlgen_helpers_code.hpp"
#include "../../../gen/schema_cpp/__meta/headers.hpp"
#include "../../../gen/schema_cpp/__factory/headers.hpp"

// for convenience
using json = nlohmann::json;

int main() {
    tl2::factory::set_all_factories();

    std::ifstream f("../data/test-functions-bytes.json");
    json data = json::parse(f);

    auto tests = data["TestReadFunction"];
    for (auto &test_data: tests) {
        std::cout << "Run [" << test_data.at("FunctionName") << ", " << test_data.at("FunctionBodyBytes") << "]: ";

        auto test_function = tl2::meta::get_item_by_name(test_data.at("FunctionName")).create_function();
        auto function_body_input = hex::parse_hex_to_bytes(test_data.at("FunctionBodyBytes"));
        auto expected_result_output = hex::parse_hex_to_bytes(test_data.at("ResultBytes"));

        basictl::tl_istream_string input1{function_body_input};
        basictl::tl_istream_string input2{expected_result_output};
        basictl::tl_ostream_string output{};

        bool read_result = test_function->read(input1);
        bool test_result = read_result;

        if (read_result) {
            test_result = test_function->read_write_result(input2, output);
            if (test_result) {
                test_result = output.get_buffer() == expected_result_output;
            }
        }

        if (test_result) {
            std::cout << "SUCCESS" << std::endl;
        } else {
            std::cout << "FAILED" << std::endl;
            std::cout << "\t\tExpected output:" << test_data.at("ResultBytes") << std::endl;
            std::cout << "\t\tActual result  :" << output.get_buffer() << std::endl;
            return 1;
        }
    }


    std::cout << std::endl;
    std::cout << "All tests are passed!" << std::endl;

    return 0;
}
