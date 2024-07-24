#pragma once
#include "../a_tlgen_helpers_code.hpp"

#include <functional>
#include <map>

namespace tl2 {
namespace meta {
    struct tl_object {
        std::function<bool(::basictl::tl_istream & )> read;
        std::function<bool(::basictl::tl_ostream & )> write;

        std::function<bool(::basictl::tl_istream & )> read_boxed;
        std::function<bool(::basictl::tl_ostream & )> write_boxed;
    };

	struct tl_function {
		tl_object object;
		
		std::function<bool(::basictl::tl_istream &, ::basictl::tl_ostream &)> read_write_result;
	};

    struct tl_item {
        uint32_t tag{};
        uint32_t annotations{};
        std::string name;

        std::function<tl2::meta::tl_object()> create_object;
		std::function<tl2::meta::tl_function()> create_function;
    };

    namespace {
        std::map<std::string, tl2::meta::tl_item> __items;

		std::function<tl_object()> missing_object_generator = []() -> tl_object {
            throw std::runtime_error("no generator for this type of object initialized");
        };
		std::function<tl_function()> missing_function_generator = []() -> tl_function {
            throw std::runtime_error("no generator for this type of function initialized");
        };
    }

    tl_item get_tl_item_by_name(std::string&& name) {
        if (__items.count(name)) {
            return __items[name];
        }
        throw std::runtime_error("no such tl (\"" + name + "\") item in system");
    }

    void set_create_object_by_name(std::string&& name, std::function<tl_object()>&& generator) {
        if (__items.count(name)) {
            __items[name].create_object = generator;
            return;
        }
        throw std::runtime_error("no such tl (\"" + name + "\") item in system");
    }

	void set_create_function_by_name(std::string&& name, std::function<tl_function()>&& generator) {
        if (__items.count(name)) {
            __items[name].create_function = generator;
            return;
        }
        throw std::runtime_error("no such tl (\"" + name + "\") item in system");
    }

    void init_tl_items() {
		__items["benchmarks.vruhash"] = tl2::meta::tl_item{.tag=0xd31bd0fd,.annotations=0x0,.name="benchmarks.vruhash",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["benchmarks.vruposition"] = tl2::meta::tl_item{.tag=0x32792c04,.annotations=0x0,.name="benchmarks.vruposition",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["benchmarks.vrutoyTopLevelContainer"] = tl2::meta::tl_item{.tag=0xfb442ca5,.annotations=0x0,.name="benchmarks.vrutoyTopLevelContainer",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["benchmarks.vrutoyTopLevelContainerWithDependency"] = tl2::meta::tl_item{.tag=0xc176008e,.annotations=0x0,.name="benchmarks.vrutoyTopLevelContainerWithDependency",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["benchmarks.vrutoytopLevelUnionBig"] = tl2::meta::tl_item{.tag=0xef556bee,.annotations=0x0,.name="benchmarks.vrutoytopLevelUnionBig",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["benchmarks.vrutoytopLevelUnionEmpty"] = tl2::meta::tl_item{.tag=0xce27c770,.annotations=0x0,.name="benchmarks.vrutoytopLevelUnionEmpty",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases_bytes.testArray"] = tl2::meta::tl_item{.tag=0x3762fb81,.annotations=0x0,.name="cases_bytes.testArray",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases_bytes.testDictAny"] = tl2::meta::tl_item{.tag=0x5a5fce57,.annotations=0x0,.name="cases_bytes.testDictAny",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases_bytes.testDictInt"] = tl2::meta::tl_item{.tag=0x453ace07,.annotations=0x0,.name="cases_bytes.testDictInt",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases_bytes.testDictString"] = tl2::meta::tl_item{.tag=0x6c04d6ce,.annotations=0x0,.name="cases_bytes.testDictString",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases_bytes.testDictStringString"] = tl2::meta::tl_item{.tag=0xad69c772,.annotations=0x0,.name="cases_bytes.testDictStringString",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases_bytes.testEnum1"] = tl2::meta::tl_item{.tag=0x58aad3f5,.annotations=0x0,.name="cases_bytes.testEnum1",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases_bytes.testEnum2"] = tl2::meta::tl_item{.tag=0x00b47add,.annotations=0x0,.name="cases_bytes.testEnum2",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases_bytes.testEnum3"] = tl2::meta::tl_item{.tag=0x81911ffa,.annotations=0x0,.name="cases_bytes.testEnum3",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases_bytes.testEnumContainer"] = tl2::meta::tl_item{.tag=0x32b92037,.annotations=0x0,.name="cases_bytes.testEnumContainer",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases_bytes.testTuple"] = tl2::meta::tl_item{.tag=0x2dd3bacf,.annotations=0x0,.name="cases_bytes.testTuple",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases_bytes.testVector"] = tl2::meta::tl_item{.tag=0x3647c8ae,.annotations=0x0,.name="cases_bytes.testVector",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.myCycle1"] = tl2::meta::tl_item{.tag=0xd3ca919d,.annotations=0x0,.name="cases.myCycle1",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.myCycle2"] = tl2::meta::tl_item{.tag=0x5444c9a2,.annotations=0x0,.name="cases.myCycle2",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.myCycle3"] = tl2::meta::tl_item{.tag=0x7624f86b,.annotations=0x0,.name="cases.myCycle3",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.replace7"] = tl2::meta::tl_item{.tag=0x6ccce4be,.annotations=0x0,.name="cases.replace7",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.replace7plus"] = tl2::meta::tl_item{.tag=0x197858f5,.annotations=0x0,.name="cases.replace7plus",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.replace7plusplus"] = tl2::meta::tl_item{.tag=0xabc39b68,.annotations=0x0,.name="cases.replace7plusplus",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.testAllPossibleFieldConfigsContainer"] = tl2::meta::tl_item{.tag=0xe3fae936,.annotations=0x0,.name="cases.testAllPossibleFieldConfigsContainer",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.testArray"] = tl2::meta::tl_item{.tag=0xa888030d,.annotations=0x0,.name="cases.testArray",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.testBeforeReadBitValidation"] = tl2::meta::tl_item{.tag=0x9b2396db,.annotations=0x0,.name="cases.testBeforeReadBitValidation",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.testDictAny"] = tl2::meta::tl_item{.tag=0xe29b8ae6,.annotations=0x0,.name="cases.testDictAny",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.testDictInt"] = tl2::meta::tl_item{.tag=0xd3877643,.annotations=0x0,.name="cases.testDictInt",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.testDictString"] = tl2::meta::tl_item{.tag=0xc463c79b,.annotations=0x0,.name="cases.testDictString",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.testEnum1"] = tl2::meta::tl_item{.tag=0x6c6c55ac,.annotations=0x0,.name="cases.testEnum1",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.testEnum2"] = tl2::meta::tl_item{.tag=0x86ea88ce,.annotations=0x0,.name="cases.testEnum2",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.testEnum3"] = tl2::meta::tl_item{.tag=0x69b83e2f,.annotations=0x0,.name="cases.testEnum3",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.testEnumContainer"] = tl2::meta::tl_item{.tag=0xcb684231,.annotations=0x0,.name="cases.testEnumContainer",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.testLocalFieldmask"] = tl2::meta::tl_item{.tag=0xf68fd3f9,.annotations=0x0,.name="cases.testLocalFieldmask",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.testMaybe"] = tl2::meta::tl_item{.tag=0xd6602613,.annotations=0x0,.name="cases.testMaybe",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.testOutFieldMaskContainer"] = tl2::meta::tl_item{.tag=0x1850ffe4,.annotations=0x0,.name="cases.testOutFieldMaskContainer",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.testRecursiveFieldMask"] = tl2::meta::tl_item{.tag=0xc58cf85e,.annotations=0x0,.name="cases.testRecursiveFieldMask",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.testTuple"] = tl2::meta::tl_item{.tag=0x4b9caf8f,.annotations=0x0,.name="cases.testTuple",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.testUnion1"] = tl2::meta::tl_item{.tag=0x4b4f09b1,.annotations=0x0,.name="cases.testUnion1",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.testUnion2"] = tl2::meta::tl_item{.tag=0x464f96c4,.annotations=0x0,.name="cases.testUnion2",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.testUnionContainer"] = tl2::meta::tl_item{.tag=0x4497a381,.annotations=0x0,.name="cases.testUnionContainer",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["cases.testVector"] = tl2::meta::tl_item{.tag=0x4975695c,.annotations=0x0,.name="cases.testVector",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["int"] = tl2::meta::tl_item{.tag=0xa8509bda,.annotations=0x0,.name="int",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["int32"] = tl2::meta::tl_item{.tag=0x7934e71f,.annotations=0x0,.name="int32",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["int64"] = tl2::meta::tl_item{.tag=0xf5609de0,.annotations=0x0,.name="int64",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["long"] = tl2::meta::tl_item{.tag=0x22076cba,.annotations=0x0,.name="long",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["string"] = tl2::meta::tl_item{.tag=0xb5286e24,.annotations=0x0,.name="string",.create_object=missing_object_generator,.create_function=missing_function_generator};
		__items["true"] = tl2::meta::tl_item{.tag=0x3fedd339,.annotations=0x0,.name="true",.create_object=missing_object_generator,.create_function=missing_function_generator};
	}
};
};
