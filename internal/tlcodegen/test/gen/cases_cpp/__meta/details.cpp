#include "../basictl/io_streams.h"

#include <map>

#include "headers.h"

namespace {
	struct tl_items {
		public:
			std::map<std::string, std::shared_ptr<tl2::meta::tl_item>> items;
			std::map<uint32_t, std::shared_ptr<tl2::meta::tl_item>> items_by_tag;
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
    auto item = items.items.find(s);
	if (item != items.items.end()) {
        return *item->second;
    }
    throw std::runtime_error("no item with such name + \"" + s + "\"");
}

tl2::meta::tl_item tl2::meta::get_item_by_tag(std::uint32_t &&tag) {
    auto item = items.items_by_tag.find(tag);
	if (item != items.items_by_tag.end()) {
        return *item->second;
    }
    throw std::runtime_error("no item with such tag + \"" + std::to_string(tag) + "\"");
}

void tl2::meta::set_create_object_by_name(std::string &&s, std::function<std::unique_ptr<tl2::meta::tl_object>()>&& gen) {
    auto item = items.items.find(s);
	if (item != items.items.end()) {
        item->second->create_object = gen;
		return;	
    }
    throw std::runtime_error("no item with such name + \"" + s + "\"");
}

void tl2::meta::set_create_function_by_name(std::string &&s, std::function<std::unique_ptr<tl2::meta::tl_function>()>&& gen) {
    auto item = items.items.find(s);
	if (item != items.items.end()) {
        item->second->create_function = gen;
		return;	
    }
    throw std::runtime_error("no item with such name + \"" + s + "\"");
}

tl_items::tl_items() {
	auto item3541815549 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xd31bd0fd,.annotations=0x0,.name="benchmarks.vruhash",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["benchmarks.vruhash"] = item3541815549;
	(this->items_by_tag)[0xd31bd0fd] = item3541815549;
	auto item846801924 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x32792c04,.annotations=0x0,.name="benchmarks.vruposition",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["benchmarks.vruposition"] = item846801924;
	(this->items_by_tag)[0x32792c04] = item846801924;
	auto item4215549093 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xfb442ca5,.annotations=0x0,.name="benchmarks.vrutoyTopLevelContainer",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["benchmarks.vrutoyTopLevelContainer"] = item4215549093;
	(this->items_by_tag)[0xfb442ca5] = item4215549093;
	auto item3245736078 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xc176008e,.annotations=0x0,.name="benchmarks.vrutoyTopLevelContainerWithDependency",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["benchmarks.vrutoyTopLevelContainerWithDependency"] = item3245736078;
	(this->items_by_tag)[0xc176008e] = item3245736078;
	auto item4015352814 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xef556bee,.annotations=0x0,.name="benchmarks.vrutoytopLevelUnionBig",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["benchmarks.vrutoytopLevelUnionBig"] = item4015352814;
	(this->items_by_tag)[0xef556bee] = item4015352814;
	auto item3458713456 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xce27c770,.annotations=0x0,.name="benchmarks.vrutoytopLevelUnionEmpty",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["benchmarks.vrutoytopLevelUnionEmpty"] = item3458713456;
	(this->items_by_tag)[0xce27c770] = item3458713456;
	auto item929233793 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x3762fb81,.annotations=0x0,.name="cases_bytes.testArray",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases_bytes.testArray"] = item929233793;
	(this->items_by_tag)[0x3762fb81] = item929233793;
	auto item1516228183 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x5a5fce57,.annotations=0x0,.name="cases_bytes.testDictAny",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases_bytes.testDictAny"] = item1516228183;
	(this->items_by_tag)[0x5a5fce57] = item1516228183;
	auto item1161481735 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x453ace07,.annotations=0x0,.name="cases_bytes.testDictInt",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases_bytes.testDictInt"] = item1161481735;
	(this->items_by_tag)[0x453ace07] = item1161481735;
	auto item1812256462 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x6c04d6ce,.annotations=0x0,.name="cases_bytes.testDictString",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases_bytes.testDictString"] = item1812256462;
	(this->items_by_tag)[0x6c04d6ce] = item1812256462;
	auto item2909390706 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xad69c772,.annotations=0x0,.name="cases_bytes.testDictStringString",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases_bytes.testDictStringString"] = item2909390706;
	(this->items_by_tag)[0xad69c772] = item2909390706;
	auto item1487590389 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x58aad3f5,.annotations=0x0,.name="cases_bytes.testEnum1",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases_bytes.testEnum1"] = item1487590389;
	(this->items_by_tag)[0x58aad3f5] = item1487590389;
	auto item11827933 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x00b47add,.annotations=0x0,.name="cases_bytes.testEnum2",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases_bytes.testEnum2"] = item11827933;
	(this->items_by_tag)[0x00b47add] = item11827933;
	auto item2173771770 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x81911ffa,.annotations=0x0,.name="cases_bytes.testEnum3",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases_bytes.testEnum3"] = item2173771770;
	(this->items_by_tag)[0x81911ffa] = item2173771770;
	auto item850993207 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x32b92037,.annotations=0x0,.name="cases_bytes.testEnumContainer",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases_bytes.testEnumContainer"] = item850993207;
	(this->items_by_tag)[0x32b92037] = item850993207;
	auto item768850639 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x2dd3bacf,.annotations=0x0,.name="cases_bytes.testTuple",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases_bytes.testTuple"] = item768850639;
	(this->items_by_tag)[0x2dd3bacf] = item768850639;
	auto item910674094 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x3647c8ae,.annotations=0x0,.name="cases_bytes.testVector",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases_bytes.testVector"] = item910674094;
	(this->items_by_tag)[0x3647c8ae] = item910674094;
	auto item3553268125 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xd3ca919d,.annotations=0x0,.name="cases.myCycle1",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.myCycle1"] = item3553268125;
	(this->items_by_tag)[0xd3ca919d] = item3553268125;
	auto item1413794210 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x5444c9a2,.annotations=0x0,.name="cases.myCycle2",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.myCycle2"] = item1413794210;
	(this->items_by_tag)[0x5444c9a2] = item1413794210;
	auto item1982134379 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x7624f86b,.annotations=0x0,.name="cases.myCycle3",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.myCycle3"] = item1982134379;
	(this->items_by_tag)[0x7624f86b] = item1982134379;
	auto item1825367230 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x6ccce4be,.annotations=0x0,.name="cases.replace7",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.replace7"] = item1825367230;
	(this->items_by_tag)[0x6ccce4be] = item1825367230;
	auto item427317493 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x197858f5,.annotations=0x0,.name="cases.replace7plus",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.replace7plus"] = item427317493;
	(this->items_by_tag)[0x197858f5] = item427317493;
	auto item2881723240 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xabc39b68,.annotations=0x0,.name="cases.replace7plusplus",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.replace7plusplus"] = item2881723240;
	(this->items_by_tag)[0xabc39b68] = item2881723240;
	auto item3824871734 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xe3fae936,.annotations=0x0,.name="cases.testAllPossibleFieldConfigsContainer",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testAllPossibleFieldConfigsContainer"] = item3824871734;
	(this->items_by_tag)[0xe3fae936] = item3824871734;
	auto item2827485965 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xa888030d,.annotations=0x0,.name="cases.testArray",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testArray"] = item2827485965;
	(this->items_by_tag)[0xa888030d] = item2827485965;
	auto item2602800859 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x9b2396db,.annotations=0x0,.name="cases.testBeforeReadBitValidation",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testBeforeReadBitValidation"] = item2602800859;
	(this->items_by_tag)[0x9b2396db] = item2602800859;
	auto item3801844454 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xe29b8ae6,.annotations=0x0,.name="cases.testDictAny",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testDictAny"] = item3801844454;
	(this->items_by_tag)[0xe29b8ae6] = item3801844454;
	auto item3548870211 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xd3877643,.annotations=0x0,.name="cases.testDictInt",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testDictInt"] = item3548870211;
	(this->items_by_tag)[0xd3877643] = item3548870211;
	auto item3294873499 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xc463c79b,.annotations=0x0,.name="cases.testDictString",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testDictString"] = item3294873499;
	(this->items_by_tag)[0xc463c79b] = item3294873499;
	auto item1819039148 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x6c6c55ac,.annotations=0x0,.name="cases.testEnum1",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testEnum1"] = item1819039148;
	(this->items_by_tag)[0x6c6c55ac] = item1819039148;
	auto item2263517390 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x86ea88ce,.annotations=0x0,.name="cases.testEnum2",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testEnum2"] = item2263517390;
	(this->items_by_tag)[0x86ea88ce] = item2263517390;
	auto item1773682223 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x69b83e2f,.annotations=0x0,.name="cases.testEnum3",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testEnum3"] = item1773682223;
	(this->items_by_tag)[0x69b83e2f] = item1773682223;
	auto item3412607537 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xcb684231,.annotations=0x0,.name="cases.testEnumContainer",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testEnumContainer"] = item3412607537;
	(this->items_by_tag)[0xcb684231] = item3412607537;
	auto item2850309150 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xa9e4441e,.annotations=0x0,.name="cases.testInplaceStructArgs",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testInplaceStructArgs"] = item2850309150;
	(this->items_by_tag)[0xa9e4441e] = item2850309150;
	auto item2862556288 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xaa9f2480,.annotations=0x0,.name="cases.testInplaceStructArgs2",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testInplaceStructArgs2"] = item2862556288;
	(this->items_by_tag)[0xaa9f2480] = item2862556288;
	auto item4136621049 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xf68fd3f9,.annotations=0x0,.name="cases.testLocalFieldmask",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testLocalFieldmask"] = item4136621049;
	(this->items_by_tag)[0xf68fd3f9] = item4136621049;
	auto item3596625427 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xd6602613,.annotations=0x0,.name="cases.testMaybe",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testMaybe"] = item3596625427;
	(this->items_by_tag)[0xd6602613] = item3596625427;
	auto item407961572 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x1850ffe4,.annotations=0x0,.name="cases.testOutFieldMaskContainer",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testOutFieldMaskContainer"] = item407961572;
	(this->items_by_tag)[0x1850ffe4] = item407961572;
	auto item3314350174 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xc58cf85e,.annotations=0x0,.name="cases.testRecursiveFieldMask",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testRecursiveFieldMask"] = item3314350174;
	(this->items_by_tag)[0xc58cf85e] = item3314350174;
	auto item1268559759 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x4b9caf8f,.annotations=0x0,.name="cases.testTuple",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testTuple"] = item1268559759;
	(this->items_by_tag)[0x4b9caf8f] = item1268559759;
	auto item1263471025 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x4b4f09b1,.annotations=0x0,.name="cases.testUnion1",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testUnion1"] = item1263471025;
	(this->items_by_tag)[0x4b4f09b1] = item1263471025;
	auto item1179621060 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x464f96c4,.annotations=0x0,.name="cases.testUnion2",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testUnion2"] = item1179621060;
	(this->items_by_tag)[0x464f96c4] = item1179621060;
	auto item1150788481 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x4497a381,.annotations=0x0,.name="cases.testUnionContainer",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testUnionContainer"] = item1150788481;
	(this->items_by_tag)[0x4497a381] = item1150788481;
	auto item1232431452 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x4975695c,.annotations=0x0,.name="cases.testVector",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["cases.testVector"] = item1232431452;
	(this->items_by_tag)[0x4975695c] = item1232431452;
	auto item2823855066 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xa8509bda,.annotations=0x0,.name="int",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["int"] = item2823855066;
	(this->items_by_tag)[0xa8509bda] = item2823855066;
	auto item2033510175 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x7934e71f,.annotations=0x0,.name="int32",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["int32"] = item2033510175;
	(this->items_by_tag)[0x7934e71f] = item2033510175;
	auto item4116749792 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xf5609de0,.annotations=0x0,.name="int64",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["int64"] = item4116749792;
	(this->items_by_tag)[0xf5609de0] = item4116749792;
	auto item570911930 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x22076cba,.annotations=0x0,.name="long",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["long"] = item570911930;
	(this->items_by_tag)[0x22076cba] = item570911930;
	auto item3039325732 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0xb5286e24,.annotations=0x0,.name="string",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["string"] = item3039325732;
	(this->items_by_tag)[0xb5286e24] = item3039325732;
	auto item1072550713 = std::shared_ptr<tl2::meta::tl_item>(new tl2::meta::tl_item{.tag=0x3fedd339,.annotations=0x0,.name="true",.create_object=no_object_generator,.create_function=no_function_generator});
	(this->items)["true"] = item1072550713;
	(this->items_by_tag)[0x3fedd339] = item1072550713;
}
