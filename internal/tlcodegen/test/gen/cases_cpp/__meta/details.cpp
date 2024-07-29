#include "../a_tlgen_helpers_code.hpp"

#include <map>

#include "headers.hpp"

namespace {
	struct tl_items {
		public:
			std::map<std::string, tl2::meta::tl_item> items;
			tl_items();
	};
    
	tl_items items;
    std::function<std::unique_ptr<tl2::meta::tl_object>()> no_object_generator = []() -> std::unique_ptr<tl2::meta::tl_object> {
        throw std::runtime_error("no generation for this type of objects");
    };
    std::function<std::unique_ptr<tl2::meta::tl_function>()> no_function_generator = []() -> std::unique_ptr<tl2::meta::tl_function> {
        throw std::runtime_error("no generation for this type of functions");
    };
}

tl2::meta::tl_item tl2::meta::get_item_by_name(std::string &&s) {
    if (items.items.count(s)) {
        return items.items[s];
    }
    throw std::runtime_error("no item with such name + \"" + s + "\"");
}

void tl2::meta::set_create_object_by_name(std::string &&s, std::function<std::unique_ptr<tl2::meta::tl_object>()>&& gen) {
    if (items.items.count(s)) {
        items.items[s].create_object = gen;
		return;
    }
    throw std::runtime_error("no item with such name + \"" + s + "\"");
}

void tl2::meta::set_create_function_by_name(std::string &&s, std::function<std::unique_ptr<tl2::meta::tl_function>()>&& gen) {
    if (items.items.count(s)) {
        items.items[s].create_function = gen;
		return;
    }
    throw std::runtime_error("no item with such name + \"" + s + "\"");
}

tl_items::tl_items() {
	(this->items)["benchmarks.vruhash"] = tl2::meta::tl_item{.tag=0xd31bd0fd,.annotations=0x0,.name="benchmarks.vruhash",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["benchmarks.vruposition"] = tl2::meta::tl_item{.tag=0x32792c04,.annotations=0x0,.name="benchmarks.vruposition",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["benchmarks.vrutoyTopLevelContainer"] = tl2::meta::tl_item{.tag=0xfb442ca5,.annotations=0x0,.name="benchmarks.vrutoyTopLevelContainer",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["benchmarks.vrutoyTopLevelContainerWithDependency"] = tl2::meta::tl_item{.tag=0xc176008e,.annotations=0x0,.name="benchmarks.vrutoyTopLevelContainerWithDependency",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["benchmarks.vrutoytopLevelUnionBig"] = tl2::meta::tl_item{.tag=0xef556bee,.annotations=0x0,.name="benchmarks.vrutoytopLevelUnionBig",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["benchmarks.vrutoytopLevelUnionEmpty"] = tl2::meta::tl_item{.tag=0xce27c770,.annotations=0x0,.name="benchmarks.vrutoytopLevelUnionEmpty",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases_bytes.testArray"] = tl2::meta::tl_item{.tag=0x3762fb81,.annotations=0x0,.name="cases_bytes.testArray",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases_bytes.testDictAny"] = tl2::meta::tl_item{.tag=0x5a5fce57,.annotations=0x0,.name="cases_bytes.testDictAny",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases_bytes.testDictInt"] = tl2::meta::tl_item{.tag=0x453ace07,.annotations=0x0,.name="cases_bytes.testDictInt",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases_bytes.testDictString"] = tl2::meta::tl_item{.tag=0x6c04d6ce,.annotations=0x0,.name="cases_bytes.testDictString",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases_bytes.testDictStringString"] = tl2::meta::tl_item{.tag=0xad69c772,.annotations=0x0,.name="cases_bytes.testDictStringString",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases_bytes.testEnum1"] = tl2::meta::tl_item{.tag=0x58aad3f5,.annotations=0x0,.name="cases_bytes.testEnum1",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases_bytes.testEnum2"] = tl2::meta::tl_item{.tag=0x00b47add,.annotations=0x0,.name="cases_bytes.testEnum2",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases_bytes.testEnum3"] = tl2::meta::tl_item{.tag=0x81911ffa,.annotations=0x0,.name="cases_bytes.testEnum3",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases_bytes.testEnumContainer"] = tl2::meta::tl_item{.tag=0x32b92037,.annotations=0x0,.name="cases_bytes.testEnumContainer",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases_bytes.testTuple"] = tl2::meta::tl_item{.tag=0x2dd3bacf,.annotations=0x0,.name="cases_bytes.testTuple",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases_bytes.testVector"] = tl2::meta::tl_item{.tag=0x3647c8ae,.annotations=0x0,.name="cases_bytes.testVector",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.myCycle1"] = tl2::meta::tl_item{.tag=0xd3ca919d,.annotations=0x0,.name="cases.myCycle1",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.myCycle2"] = tl2::meta::tl_item{.tag=0x5444c9a2,.annotations=0x0,.name="cases.myCycle2",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.myCycle3"] = tl2::meta::tl_item{.tag=0x7624f86b,.annotations=0x0,.name="cases.myCycle3",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.replace7"] = tl2::meta::tl_item{.tag=0x6ccce4be,.annotations=0x0,.name="cases.replace7",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.replace7plus"] = tl2::meta::tl_item{.tag=0x197858f5,.annotations=0x0,.name="cases.replace7plus",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.replace7plusplus"] = tl2::meta::tl_item{.tag=0xabc39b68,.annotations=0x0,.name="cases.replace7plusplus",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.testAllPossibleFieldConfigsContainer"] = tl2::meta::tl_item{.tag=0xe3fae936,.annotations=0x0,.name="cases.testAllPossibleFieldConfigsContainer",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.testArray"] = tl2::meta::tl_item{.tag=0xa888030d,.annotations=0x0,.name="cases.testArray",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.testBeforeReadBitValidation"] = tl2::meta::tl_item{.tag=0x9b2396db,.annotations=0x0,.name="cases.testBeforeReadBitValidation",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.testDictAny"] = tl2::meta::tl_item{.tag=0xe29b8ae6,.annotations=0x0,.name="cases.testDictAny",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.testDictInt"] = tl2::meta::tl_item{.tag=0xd3877643,.annotations=0x0,.name="cases.testDictInt",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.testDictString"] = tl2::meta::tl_item{.tag=0xc463c79b,.annotations=0x0,.name="cases.testDictString",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.testEnum1"] = tl2::meta::tl_item{.tag=0x6c6c55ac,.annotations=0x0,.name="cases.testEnum1",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.testEnum2"] = tl2::meta::tl_item{.tag=0x86ea88ce,.annotations=0x0,.name="cases.testEnum2",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.testEnum3"] = tl2::meta::tl_item{.tag=0x69b83e2f,.annotations=0x0,.name="cases.testEnum3",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.testEnumContainer"] = tl2::meta::tl_item{.tag=0xcb684231,.annotations=0x0,.name="cases.testEnumContainer",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.testLocalFieldmask"] = tl2::meta::tl_item{.tag=0xf68fd3f9,.annotations=0x0,.name="cases.testLocalFieldmask",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.testMaybe"] = tl2::meta::tl_item{.tag=0xd6602613,.annotations=0x0,.name="cases.testMaybe",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.testOutFieldMaskContainer"] = tl2::meta::tl_item{.tag=0x1850ffe4,.annotations=0x0,.name="cases.testOutFieldMaskContainer",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.testRecursiveFieldMask"] = tl2::meta::tl_item{.tag=0xc58cf85e,.annotations=0x0,.name="cases.testRecursiveFieldMask",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.testTuple"] = tl2::meta::tl_item{.tag=0x4b9caf8f,.annotations=0x0,.name="cases.testTuple",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.testUnion1"] = tl2::meta::tl_item{.tag=0x4b4f09b1,.annotations=0x0,.name="cases.testUnion1",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.testUnion2"] = tl2::meta::tl_item{.tag=0x464f96c4,.annotations=0x0,.name="cases.testUnion2",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.testUnionContainer"] = tl2::meta::tl_item{.tag=0x4497a381,.annotations=0x0,.name="cases.testUnionContainer",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["cases.testVector"] = tl2::meta::tl_item{.tag=0x4975695c,.annotations=0x0,.name="cases.testVector",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["int"] = tl2::meta::tl_item{.tag=0xa8509bda,.annotations=0x0,.name="int",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["int32"] = tl2::meta::tl_item{.tag=0x7934e71f,.annotations=0x0,.name="int32",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["int64"] = tl2::meta::tl_item{.tag=0xf5609de0,.annotations=0x0,.name="int64",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["long"] = tl2::meta::tl_item{.tag=0x22076cba,.annotations=0x0,.name="long",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["string"] = tl2::meta::tl_item{.tag=0xb5286e24,.annotations=0x0,.name="string",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["true"] = tl2::meta::tl_item{.tag=0x3fedd339,.annotations=0x0,.name="true",.create_object=no_object_generator,.create_function=no_function_generator};
}
