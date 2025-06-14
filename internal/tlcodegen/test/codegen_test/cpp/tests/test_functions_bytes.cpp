#include <iostream>
#include <fstream>
#include <bitset>

#include "../utils/hex.h"
#include "../dependencies/json.hpp"

#include "../../../gen/schema_cpp/basictl/io_streams.h"
#include "../../../gen/schema_cpp/basictl/impl/string_io.h"

#include "../../../gen/schema_cpp/__meta/headers.h"
#include "../../../gen/schema_cpp/__factory/headers.h"

// for convenience
using json = nlohmann::json;

int main() {
    tlgen::factory::set_all_factories();

    std::ifstream f("../data/test-functions-bytes.json");
    json data = json::parse(f);

    auto tests = data["TestReadFunction"];
    for (auto &test_data: tests) {
        std::cout << "Run [" << test_data.at("FunctionName") << ", " << test_data.at("FunctionBodyBytes") << "]: ";

        auto test_function = tlgen::meta::get_item_by_name(test_data.at("FunctionName")).value().create_function();
        auto function_body_input = hex::parse_hex_to_bytes(test_data.at("FunctionBodyBytes"));
        auto expected_result_output = hex::parse_hex_to_bytes(test_data.at("ResultBytes"));

        tlgen::basictl::tl_istream_string input1_connector{function_body_input};
        tlgen::basictl::tl_istream input1{input1_connector};

        tlgen::basictl::tl_istream_string input2_connector{expected_result_output};
        tlgen::basictl::tl_istream input2{input2_connector};

        std::string out_string;
        tlgen::basictl::tl_ostream_string output_connector{out_string};
        tlgen::basictl::tl_ostream output{output_connector};


        bool read_result = test_function->read(input1);
        bool test_result = read_result;

        if (read_result) {
            test_result = test_function->read_write_result(input2, output);
            if (test_result) {
                auto result = output_connector.used_buffer();
                std::string string_result(reinterpret_cast<char *>(result.data()), result.size());
                test_result = string_result == expected_result_output;
            }
        }

        if (test_result) {
            std::cout << "SUCCESS" << std::endl;
        } else {
            auto result = output_connector.used_buffer();
            std::string string_result(reinterpret_cast<char *>(result.data()), result.size());
            std::cout << "FAILED" << std::endl;
            std::cout << "\t\tExpected output:" << test_data.at("ResultBytes") << std::endl;
            std::cout << "\t\tActual result  :" << string_result << std::endl;
            return 1;
        }
    }


    std::cout << std::endl;
    std::cout << "All tests are passed!" << std::endl;

    return 0;
}
