#include <iostream>
#include <fstream>
#include <bitset>
#include "../dependencies/json.hpp"
#include "../utils/hex.h"

#include "../../../gen/cases_cpp/basictl/io_streams.h"
#include "../../../gen/cases_cpp/basictl/io_throwable_streams.h"
#include "../../../gen/cases_cpp/basictl/impl/string_io.h"

#include "../../../gen/cases_cpp/__meta/headers.h"
#include "../../../gen/cases_cpp/__factory/headers.h"

// for convenience
using json = nlohmann::json;

bool test_case(std::string type_name, std::string& testing_bytes);
bool test_case_throwable(std::string type_name, std::string& testing_bytes);

int main() {
    tl2::factory::set_all_factories();

    std::ifstream f("../data/test-objects-bytes.json");
    json data = json::parse(f);

    auto tests = data["Tests"];
    for (auto& [test_name, test_data]: tests.items()) {
        std::cout << "Run [" << test_name << "]:" << std::endl;
        for (auto& test_data_input: test_data["Successes"]) {
            std::cout << "\tTestData [" << test_data_input.at("Bytes") << "]: ";

            std::string type_name = test_data.at("TestingType");
            std::string testing_bytes = test_data_input.at("Bytes");

            auto success = test_case(type_name, testing_bytes);
            if (!success) {
                std::cout << std::endl;
                return 1;
            }

            std::cout << ", ";

            auto success_throwable = test_case_throwable(type_name, testing_bytes);
            if (!success_throwable) {
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
    auto test_object = tl2::meta::get_item_by_name(std::move(type_name)).create_object();
    auto expected_output = hex::parse_hex_to_bytes(testing_bytes);

    basictl::tl_istream_string input_connector{expected_output};
    basictl::tl_istream input{input_connector};

    std::string out_string;
    basictl::tl_ostream_string output_connector{out_string};
    basictl::tl_ostream output{output_connector};

    bool read_result = test_object->read(input);
    bool write_result = test_object->write(output);

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


bool test_case_throwable(std::string type_name, std::string& testing_bytes) {
    auto test_object = tl2::meta::get_item_by_name(std::move(type_name)).create_object();
    auto expected_output = hex::parse_hex_to_bytes(testing_bytes);

    basictl::tl_istream_string input_connector{expected_output};
    basictl::tl_throwable_istream input{input_connector};

    std::string out_string;
    basictl::tl_ostream_string output_connector{out_string};
    basictl::tl_throwable_ostream output{output_connector};

    try {
        test_object->read_or_throw(input);
        test_object->write_or_throw(output);

        std::cout << "SUCCESS (THROWABLE)";
        return true;
    } catch (std::exception& e) {
        auto result = output_connector.used_buffer();
        std::string string_result(reinterpret_cast<char *>(result.data()), result.size());
        std::cout << "FAILED (THROWABLE)" << std::endl;
        std::cout << "\t\tExpected output:" << expected_output << std::endl;
        std::cout << "\t\tActual result  :" << string_result << std::endl;
        std::cout << "\t\tError reason   :" << e.what() << std::endl;
        return false;
    }
}

