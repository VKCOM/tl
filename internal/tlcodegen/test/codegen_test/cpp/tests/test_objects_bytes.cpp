#include <iostream>
#include <fstream>
#include <bitset>
#include "../dependencies/json.hpp"
#include "../utils/hex.h"
#include "../../../gen/cases_cpp/a_tlgen_helpers_code.hpp"

#include "../../../gen/cases_cpp/__meta/meta.hpp"
#include "../../../gen/cases_cpp/__factory/factory.hpp"

// for convenience
using json = nlohmann::json;

int main() {
    tl2::meta::init_tl_items();
    tl2::factory::init_tl_create_objects();

    std::ifstream f("../data/test-objects-bytes.json");
    json data = json::parse(f);

    auto tests = data["Tests"];
    for (auto& [test_name, test_data]: tests.items()) {
        std::cout << "Run [" << test_name << "]:" << std::endl;
        for (auto& test_data_input: test_data["Successes"]) {
            std::cout << "\tTestData [" << test_data_input.at("Bytes") << "]: ";

            auto test_object = tl2::meta::get_tl_item_by_name(test_data.at("TestingType")).create_object();
            auto expected_output = hex::parse_hex_to_bytes(test_data_input.at("Bytes"));

            basictl::tl_istream_string input{expected_output};
            basictl::tl_ostream_string output{};

            bool read_result = test_object.read(input);
            bool write_result = test_object.write(output);

            bool test_result = write_result && read_result;
            if (test_result) {
                test_result = output.get_buffer() == expected_output;
            }
            if (test_result) {
                std::cout << "SUCCESS" << std::endl;
            } else {
                std::cout << "FAILED" << std::endl;
                std::cout << "\t\tWrite result   :" << write_result << std::endl;
                std::cout << "\t\tExpected output:" << expected_output << std::endl;
                std::cout << "\t\tActual result  :" << output.get_buffer() << std::endl;
                return 1;
            }
        }
    }

    std::cout << std::endl;
    std::cout << "All tests are passed!" << std::endl;

    return 0;
}
