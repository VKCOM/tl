#include <iostream>
#include <fstream>
#include <bitset>
#include "../dependencies/json.hpp"
#include "../utils/hex.h"

#include "../../../gen/schema_cpp/basics/basictl.h"
#include "../../../gen/schema_cpp/meta/headers.h"
#include "../../../gen/schema_cpp/factory/headers.h"
#include "../../../gen/schema_cpp/basics/string_io.h"

// for convenience
using json = nlohmann::json;

int main() {
    tl2::factory::set_all_factories();

    std::ifstream f("../data/test-objects-json-and-bytes.json");
    json data = json::parse(f);

    auto tests = data["TestsCpp"];
    for (auto& [test_name, test_data]: tests.items()) {
        std::cout << "Run [" << test_name << "]:" << std::endl;
        for (auto& test_data_input: test_data["Successes"]) {
            std::cout << "\tTestData [" << test_data_input.at("DataAsBytes") << "]: ";

            auto test_object = tl2::meta::get_item_by_name(test_data.at("TestingType")).create_object();
            auto bytes_input = hex::parse_hex_to_bytes(test_data_input.at("DataAsBytes"));
            std::string expected_json_output = test_data_input.at("DataAsJson");

            auto string_input = basictl::tl_istream_string{bytes_input};
            basictl::tl_istream input = basictl::tl_istream(&string_input);

            std::stringstream output;

            bool read_result = test_object->read(input);
            bool write_result = test_object->write_json(output);

            bool test_result = write_result && read_result;
            if (test_result) {
                test_result = output.str() == expected_json_output;
            }
            if (test_result) {
                std::cout << "SUCCESS" << std::endl;
            } else {
                std::cout << "FAILED" << std::endl;
                std::cout << "\t\tWrite result        : " << write_result << std::endl;
                std::cout << "\t\tExpected json output: " << expected_json_output << std::endl;
                std::cout << "\t\tActual json result  : " << output.str() << std::endl;
                return 1;
            }
        }
    }

    std::cout << std::endl;
    std::cout << "All tests are passed!" << std::endl;

    return 0;
}