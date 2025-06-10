#include <iostream>
#include <fstream>
#include <bitset>
#include "../dependencies/json.hpp"
#include "../utils/hex.h"

#include "../../../gen/goldmaster_cpp/basictl/io_streams.h"
#include "../../../gen/goldmaster_cpp/basictl/io_throwable_streams.h"
#include "../../../gen/goldmaster_cpp/basictl/impl/string_io.h"

#include "../../../gen/goldmaster_cpp/__meta/headers.h"
#include "../../../gen/goldmaster_cpp/__factory/headers.h"

// for convenience
using json = nlohmann::json;

bool test_case(std::string type_name, std::string& testing_bytes);
bool test_case_throwable(std::string type_name, std::string& testing_bytes);

int main() {
    tlgen::factory::set_all_factories();

    std::ifstream f("../data/test-stress-data-goldmaster.json");
    json data = json::parse(f);

    auto tests = data["Tests"];
    for (auto& [test_name, test_data]: tests.items()) {
        std::cout << "Run [" << test_name << "]:" << std::endl;
        for (auto& test_data_input: test_data["Successes"]) {
            std::string type_name = test_data.at("TestingType");
            std::string testing_bytes = test_data_input.at("Bytes");

            std::cout << "\tTestData [" << testing_bytes << "]: ";

            auto success = test_case(type_name, testing_bytes);
            if (!success) {
                std::cout << std::endl;
                return 1;
            }

            std::cout << std::endl;
        }
    }

    std::cout << std::endl;
    std::cout << "All tests are passed!" << std::endl;

    return 0;
}


bool test_case(std::string type_name, std::string& testing_bytes) {
    auto item = tlgen::meta::get_item_by_name(std::move(type_name));
    if (!item || !item.value().has_create_object) {
        return true;
    }
    auto test_object = item.value().create_object();
    auto expected_output = hex::parse_hex_to_bytes(testing_bytes);

    tlgen::basictl::tl_istream_string input_connector{expected_output};
    tlgen::basictl::tl_istream input{input_connector};

    std::string out_string;
    tlgen::basictl::tl_ostream_string output_connector{out_string};
    tlgen::basictl::tl_ostream output{output_connector};

    bool read_result = test_object->read_boxed(input);
    bool write_result = test_object->write_boxed(output);

    bool test_result = write_result && read_result;
    auto result = output_connector.used_buffer();
    std::string string_result(reinterpret_cast<char *>(result.data()), result.size());

    if (test_result) {
        test_result = string_result == expected_output;
    }
    if (test_result) {
        std::cout << "SUCCESS (NON-THROWABLE)";
    } else {
        std::cout << "FAILED (NON-THROWABLE)" << std::endl;
        std::cout << "\t\tWrite result   :" << write_result << std::endl;
        std::cout << "\t\tExpected output:" << expected_output << std::endl;
        std::cout << "\t\tActual result  :" << string_result << std::endl;
        return false;
    }
    return true;
}

