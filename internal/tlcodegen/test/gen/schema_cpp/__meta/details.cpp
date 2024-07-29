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
	(this->items)["antispam.getPattern"] = tl2::meta::tl_item{.tag=0x3de14136,.annotations=0x1,.name="antispam.getPattern",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["antispam.patternFound"] = tl2::meta::tl_item{.tag=0xa7688492,.annotations=0x0,.name="antispam.patternFound",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["antispam.patternNotFound"] = tl2::meta::tl_item{.tag=0x2c22e225,.annotations=0x0,.name="antispam.patternNotFound",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["benchObject"] = tl2::meta::tl_item{.tag=0xb697e865,.annotations=0x0,.name="benchObject",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["boolStat"] = tl2::meta::tl_item{.tag=0x92cbcbfa,.annotations=0x0,.name="boolStat",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["boxedArray"] = tl2::meta::tl_item{.tag=0x95dcc8b7,.annotations=0x1,.name="boxedArray",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["boxedInt"] = tl2::meta::tl_item{.tag=0x5688ebaf,.annotations=0x1,.name="boxedInt",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["boxedString"] = tl2::meta::tl_item{.tag=0x548994db,.annotations=0x1,.name="boxedString",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["boxedTuple"] = tl2::meta::tl_item{.tag=0x30c9d533,.annotations=0x1,.name="boxedTuple",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["boxedTupleSlice1"] = tl2::meta::tl_item{.tag=0x25230d40,.annotations=0x1,.name="boxedTupleSlice1",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["boxedTupleSlice2"] = tl2::meta::tl_item{.tag=0x1cdf4705,.annotations=0x1,.name="boxedTupleSlice2",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["boxedTupleSlice3"] = tl2::meta::tl_item{.tag=0xa19b8106,.annotations=0x1,.name="boxedTupleSlice3",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["boxedVector32"] = tl2::meta::tl_item{.tag=0xbbadef07,.annotations=0x1,.name="boxedVector32",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["boxedVector32BoxedElem"] = tl2::meta::tl_item{.tag=0x591cecd4,.annotations=0x1,.name="boxedVector32BoxedElem",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["boxedVector64"] = tl2::meta::tl_item{.tag=0x83659ba8,.annotations=0x1,.name="boxedVector64",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["double"] = tl2::meta::tl_item{.tag=0x2210c154,.annotations=0x0,.name="double",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["fieldConflict1"] = tl2::meta::tl_item{.tag=0xf314bd09,.annotations=0x0,.name="fieldConflict1",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["fieldConflict2"] = tl2::meta::tl_item{.tag=0x1bba76b8,.annotations=0x0,.name="fieldConflict2",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["fieldConflict3"] = tl2::meta::tl_item{.tag=0x2cf6e157,.annotations=0x0,.name="fieldConflict3",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["fieldConflict4"] = tl2::meta::tl_item{.tag=0xd93c186a,.annotations=0x0,.name="fieldConflict4",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["float"] = tl2::meta::tl_item{.tag=0x824dab22,.annotations=0x0,.name="float",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["get_arrays"] = tl2::meta::tl_item{.tag=0x90658cdb,.annotations=0x1,.name="get_arrays",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["getDouble"] = tl2::meta::tl_item{.tag=0x39711d7b,.annotations=0x1,.name="getDouble",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["getFloat"] = tl2::meta::tl_item{.tag=0x25a7bc68,.annotations=0x1,.name="getFloat",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["getMaybeIface"] = tl2::meta::tl_item{.tag=0x6b055ae4,.annotations=0x1,.name="getMaybeIface",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["getMyDictOfInt"] = tl2::meta::tl_item{.tag=0x166f962c,.annotations=0x1,.name="getMyDictOfInt",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["getMyDouble"] = tl2::meta::tl_item{.tag=0xb660ad10,.annotations=0x1,.name="getMyDouble",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["getMyValue"] = tl2::meta::tl_item{.tag=0xb3df27fe,.annotations=0x1,.name="getMyValue",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["getNonOptNat"] = tl2::meta::tl_item{.tag=0x67665961,.annotations=0x1,.name="getNonOptNat",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["getStats"] = tl2::meta::tl_item{.tag=0xbaa6da35,.annotations=0x1,.name="getStats",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["int"] = tl2::meta::tl_item{.tag=0xa8509bda,.annotations=0x0,.name="int",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["integer"] = tl2::meta::tl_item{.tag=0x7e194796,.annotations=0x0,.name="integer",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["issue3498"] = tl2::meta::tl_item{.tag=0xf54b7b0a,.annotations=0x0,.name="issue3498",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["long"] = tl2::meta::tl_item{.tag=0x22076cba,.annotations=0x0,.name="long",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["myAnonMcValue"] = tl2::meta::tl_item{.tag=0x569310db,.annotations=0x0,.name="myAnonMcValue",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["myBoxedArray"] = tl2::meta::tl_item{.tag=0x288f64f0,.annotations=0x0,.name="myBoxedArray",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["myBoxedTupleSlice"] = tl2::meta::tl_item{.tag=0x25d1a1be,.annotations=0x0,.name="myBoxedTupleSlice",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["myBoxedVectorSlice"] = tl2::meta::tl_item{.tag=0x57d164bb,.annotations=0x0,.name="myBoxedVectorSlice",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["myDictOfInt"] = tl2::meta::tl_item{.tag=0xb8019a3d,.annotations=0x0,.name="myDictOfInt",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["myDouble"] = tl2::meta::tl_item{.tag=0x90a6c726,.annotations=0x0,.name="myDouble",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["myInt"] = tl2::meta::tl_item{.tag=0xc12375b7,.annotations=0x0,.name="myInt",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["myMaybe1"] = tl2::meta::tl_item{.tag=0x32c541fe,.annotations=0x0,.name="myMaybe1",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["myMaybe2"] = tl2::meta::tl_item{.tag=0xef6d355c,.annotations=0x0,.name="myMaybe2",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["myMcValue"] = tl2::meta::tl_item{.tag=0xe2ffd978,.annotations=0x0,.name="myMcValue",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["myMcValueTuple"] = tl2::meta::tl_item{.tag=0x1287d116,.annotations=0x0,.name="myMcValueTuple",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["myMcValueVector"] = tl2::meta::tl_item{.tag=0x761d6d58,.annotations=0x0,.name="myMcValueVector",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["myString"] = tl2::meta::tl_item{.tag=0xc8bfa969,.annotations=0x0,.name="myString",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["myTwoDicts"] = tl2::meta::tl_item{.tag=0xa859581d,.annotations=0x0,.name="myTwoDicts",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["nonOptNat"] = tl2::meta::tl_item{.tag=0x45366605,.annotations=0x0,.name="nonOptNat",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["pkg2.foo"] = tl2::meta::tl_item{.tag=0xe144703d,.annotations=0x0,.name="pkg2.foo",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["pkg2.t1"] = tl2::meta::tl_item{.tag=0x638206ec,.annotations=0x0,.name="pkg2.t1",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["pkg2.t2"] = tl2::meta::tl_item{.tag=0xd6e5af9c,.annotations=0x0,.name="pkg2.t2",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["rpcInvokeReqExtra"] = tl2::meta::tl_item{.tag=0xf3ef81a9,.annotations=0x0,.name="rpcInvokeReqExtra",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.add"] = tl2::meta::tl_item{.tag=0x481df8be,.annotations=0x1,.name="service1.add",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.addOrGet"] = tl2::meta::tl_item{.tag=0x6a42faad,.annotations=0x1,.name="service1.addOrGet",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.addOrIncr"] = tl2::meta::tl_item{.tag=0x90c4b402,.annotations=0x1,.name="service1.addOrIncr",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.append"] = tl2::meta::tl_item{.tag=0x04dec671,.annotations=0x1,.name="service1.append",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.cas"] = tl2::meta::tl_item{.tag=0x51851964,.annotations=0x1,.name="service1.cas",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.decr"] = tl2::meta::tl_item{.tag=0xeb179ce7,.annotations=0x1,.name="service1.decr",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.delete"] = tl2::meta::tl_item{.tag=0x83277767,.annotations=0x1,.name="service1.delete",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.disableExpiration"] = tl2::meta::tl_item{.tag=0xf1c39c2d,.annotations=0x1,.name="service1.disableExpiration",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.disableKeysStat"] = tl2::meta::tl_item{.tag=0x79d6160f,.annotations=0x1,.name="service1.disableKeysStat",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.enableExpiration"] = tl2::meta::tl_item{.tag=0x2b51ad67,.annotations=0x1,.name="service1.enableExpiration",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.enableKeysStat"] = tl2::meta::tl_item{.tag=0x29a7090e,.annotations=0x1,.name="service1.enableKeysStat",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.exists"] = tl2::meta::tl_item{.tag=0xe0284c9e,.annotations=0x1,.name="service1.exists",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.get"] = tl2::meta::tl_item{.tag=0x29099b19,.annotations=0x1,.name="service1.get",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.getExpireTime"] = tl2::meta::tl_item{.tag=0x5a731070,.annotations=0x1,.name="service1.getExpireTime",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.getKeysStat"] = tl2::meta::tl_item{.tag=0x06cecd58,.annotations=0x1,.name="service1.getKeysStat",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.getKeysStatPeriods"] = tl2::meta::tl_item{.tag=0x8cdf39e3,.annotations=0x1,.name="service1.getKeysStatPeriods",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.getWildcard"] = tl2::meta::tl_item{.tag=0x2f2abf13,.annotations=0x1,.name="service1.getWildcard",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.getWildcardDict"] = tl2::meta::tl_item{.tag=0x72bbc81b,.annotations=0x1,.name="service1.getWildcardDict",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.getWildcardList"] = tl2::meta::tl_item{.tag=0x56b6ead4,.annotations=0x1,.name="service1.getWildcardList",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.getWildcardWithFlags"] = tl2::meta::tl_item{.tag=0x5f6a1f78,.annotations=0x1,.name="service1.getWildcardWithFlags",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.incr"] = tl2::meta::tl_item{.tag=0x0f96b56e,.annotations=0x1,.name="service1.incr",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.keysStat"] = tl2::meta::tl_item{.tag=0xf0f6bc68,.annotations=0x0,.name="service1.keysStat",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.longvalue"] = tl2::meta::tl_item{.tag=0x082e0945,.annotations=0x0,.name="service1.longvalue",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.longvalueWithTime"] = tl2::meta::tl_item{.tag=0xa04606ec,.annotations=0x0,.name="service1.longvalueWithTime",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.not_found"] = tl2::meta::tl_item{.tag=0x1d670b96,.annotations=0x0,.name="service1.not_found",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.replace"] = tl2::meta::tl_item{.tag=0x7f2c447d,.annotations=0x1,.name="service1.replace",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.replaceOrIncr"] = tl2::meta::tl_item{.tag=0x9d1bdcfd,.annotations=0x1,.name="service1.replaceOrIncr",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.set"] = tl2::meta::tl_item{.tag=0x05ae5f66,.annotations=0x1,.name="service1.set",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.setOrIncr"] = tl2::meta::tl_item{.tag=0x772e390d,.annotations=0x1,.name="service1.setOrIncr",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.strvalue"] = tl2::meta::tl_item{.tag=0x5faa0c52,.annotations=0x0,.name="service1.strvalue",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.strvalueWithTime"] = tl2::meta::tl_item{.tag=0x98b1a484,.annotations=0x0,.name="service1.strvalueWithTime",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service1.touch"] = tl2::meta::tl_item{.tag=0xb737aa03,.annotations=0x1,.name="service1.touch",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service2.addOrIncrMany"] = tl2::meta::tl_item{.tag=0x5aa52489,.annotations=0x2,.name="service2.addOrIncrMany",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service2.set"] = tl2::meta::tl_item{.tag=0x0d31f63d,.annotations=0x4,.name="service2.set",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service2.setObjectTtl"] = tl2::meta::tl_item{.tag=0x6f98f025,.annotations=0x4,.name="service2.setObjectTtl",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service3.createProduct"] = tl2::meta::tl_item{.tag=0xb7d92bd9,.annotations=0x1,.name="service3.createProduct",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service3.deleteAllProducts"] = tl2::meta::tl_item{.tag=0x4494acc2,.annotations=0x1,.name="service3.deleteAllProducts",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service3.deleteGroupedProducts"] = tl2::meta::tl_item{.tag=0xe468e614,.annotations=0x1,.name="service3.deleteGroupedProducts",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service3.deleteProduct"] = tl2::meta::tl_item{.tag=0x6867e707,.annotations=0x1,.name="service3.deleteProduct",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service3.getLastVisitTimestamp"] = tl2::meta::tl_item{.tag=0x9a4c788d,.annotations=0x1,.name="service3.getLastVisitTimestamp",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service3.getLimits"] = tl2::meta::tl_item{.tag=0xeb399467,.annotations=0x1,.name="service3.getLimits",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service3.getProductStats"] = tl2::meta::tl_item{.tag=0x261f6898,.annotations=0x1,.name="service3.getProductStats",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service3.getProducts"] = tl2::meta::tl_item{.tag=0xeb306233,.annotations=0x1,.name="service3.getProducts",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service3.getScheduledProducts"] = tl2::meta::tl_item{.tag=0xf53ad7bd,.annotations=0x1,.name="service3.getScheduledProducts",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service3.groupCountLimit"] = tl2::meta::tl_item{.tag=0x8c04ea7f,.annotations=0x0,.name="service3.groupCountLimit",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service3.groupSizeLimit"] = tl2::meta::tl_item{.tag=0x90e59396,.annotations=0x0,.name="service3.groupSizeLimit",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service3.limits"] = tl2::meta::tl_item{.tag=0x80ee61ca,.annotations=0x0,.name="service3.limits",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service3.productStatsOld"] = tl2::meta::tl_item{.tag=0x6319810b,.annotations=0x0,.name="service3.productStatsOld",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service3.restoreAllProducts"] = tl2::meta::tl_item{.tag=0x4d839ed0,.annotations=0x1,.name="service3.restoreAllProducts",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service3.restoreGroupedProducts"] = tl2::meta::tl_item{.tag=0x1f17bfac,.annotations=0x1,.name="service3.restoreGroupedProducts",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service3.restoreProduct"] = tl2::meta::tl_item{.tag=0x6170d515,.annotations=0x1,.name="service3.restoreProduct",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service3.setLastVisitTimestamp"] = tl2::meta::tl_item{.tag=0x7909b020,.annotations=0x1,.name="service3.setLastVisitTimestamp",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service3.setLimits"] = tl2::meta::tl_item{.tag=0x3ad5c19c,.annotations=0x1,.name="service3.setLimits",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service4.modifiedNewsEntry"] = tl2::meta::tl_item{.tag=0xda19832a,.annotations=0x0,.name="service4.modifiedNewsEntry",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service4.object"] = tl2::meta::tl_item{.tag=0xa6eeca4f,.annotations=0x0,.name="service4.object",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service5.emptyOutput"] = tl2::meta::tl_item{.tag=0x11e46879,.annotations=0x0,.name="service5.emptyOutput",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service5.insert"] = tl2::meta::tl_item{.tag=0xc911ee2c,.annotations=0x1,.name="service5.insert",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service5.params"] = tl2::meta::tl_item{.tag=0x12ae5cb5,.annotations=0x0,.name="service5.params",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service5.performQuery"] = tl2::meta::tl_item{.tag=0x019d80a5,.annotations=0x1,.name="service5.performQuery",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service5.query"] = tl2::meta::tl_item{.tag=0xb3b62513,.annotations=0x1,.name="service5.query",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service5.stringOutput"] = tl2::meta::tl_item{.tag=0x179e9863,.annotations=0x0,.name="service5.stringOutput",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service6.error"] = tl2::meta::tl_item{.tag=0x738553ef,.annotations=0x0,.name="service6.error",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service6.findResultRow"] = tl2::meta::tl_item{.tag=0xbd3946e3,.annotations=0x0,.name="service6.findResultRow",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service6.findWithBoundsResult"] = tl2::meta::tl_item{.tag=0x3ded850a,.annotations=0x0,.name="service6.findWithBoundsResult",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service6.multiFind"] = tl2::meta::tl_item{.tag=0xe62178d8,.annotations=0x1,.name="service6.multiFind",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["service6.multiFindWithBounds"] = tl2::meta::tl_item{.tag=0x84b168cf,.annotations=0x1,.name="service6.multiFindWithBounds",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["statOne"] = tl2::meta::tl_item{.tag=0x74b0604b,.annotations=0x0,.name="statOne",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["string"] = tl2::meta::tl_item{.tag=0xb5286e24,.annotations=0x0,.name="string",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["tasks.addTask"] = tl2::meta::tl_item{.tag=0x2ca073d5,.annotations=0x1,.name="tasks.addTask",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["tasks.cronTask"] = tl2::meta::tl_item{.tag=0xc90cf28a,.annotations=0x0,.name="tasks.cronTask",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["tasks.cronTaskWithId"] = tl2::meta::tl_item{.tag=0x3a958001,.annotations=0x0,.name="tasks.cronTaskWithId",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["tasks.cronTime"] = tl2::meta::tl_item{.tag=0xd4177d7f,.annotations=0x0,.name="tasks.cronTime",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["tasks.getAnyTask"] = tl2::meta::tl_item{.tag=0x4a9c7dbb,.annotations=0x1,.name="tasks.getAnyTask",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["tasks.getQueueSize"] = tl2::meta::tl_item{.tag=0xd8fcda03,.annotations=0x1,.name="tasks.getQueueSize",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["tasks.getQueueTypes"] = tl2::meta::tl_item{.tag=0x5434457a,.annotations=0x1,.name="tasks.getQueueTypes",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["tasks.getTaskFromQueue"] = tl2::meta::tl_item{.tag=0x6a52b698,.annotations=0x1,.name="tasks.getTaskFromQueue",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["tasks.queueTypeInfo"] = tl2::meta::tl_item{.tag=0x38d38d3e,.annotations=0x0,.name="tasks.queueTypeInfo",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["tasks.queueTypeSettings"] = tl2::meta::tl_item{.tag=0x561fbc09,.annotations=0x0,.name="tasks.queueTypeSettings",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["tasks.queueTypeStats"] = tl2::meta::tl_item{.tag=0xe1b785f2,.annotations=0x0,.name="tasks.queueTypeStats",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["tasks.task"] = tl2::meta::tl_item{.tag=0x7c23bc2c,.annotations=0x0,.name="tasks.task",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["tasks.taskInfo"] = tl2::meta::tl_item{.tag=0x06f0c6a6,.annotations=0x0,.name="tasks.taskInfo",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["tasks.taskStatusInProgress"] = tl2::meta::tl_item{.tag=0x06ef70e7,.annotations=0x0,.name="tasks.taskStatusInProgress",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["tasks.taskStatusNotCurrentlyInEngine"] = tl2::meta::tl_item{.tag=0xb207caaa,.annotations=0x0,.name="tasks.taskStatusNotCurrentlyInEngine",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["tasks.taskStatusScheduled"] = tl2::meta::tl_item{.tag=0x0aca80a9,.annotations=0x0,.name="tasks.taskStatusScheduled",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["tasks.taskStatusWaiting"] = tl2::meta::tl_item{.tag=0x16739c2c,.annotations=0x0,.name="tasks.taskStatusWaiting",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["tree_stats.objectLimitValueDouble"] = tl2::meta::tl_item{.tag=0x5dfb8816,.annotations=0x0,.name="tree_stats.objectLimitValueDouble",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["tree_stats.objectLimitValueLong"] = tl2::meta::tl_item{.tag=0x73111993,.annotations=0x0,.name="tree_stats.objectLimitValueLong",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["true"] = tl2::meta::tl_item{.tag=0x3fedd339,.annotations=0x0,.name="true",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["unique.get"] = tl2::meta::tl_item{.tag=0xce89bbf2,.annotations=0x1,.name="unique.get",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["unique.stringToInt"] = tl2::meta::tl_item{.tag=0x0f766c35,.annotations=0x1,.name="unique.stringToInt",.create_object=no_object_generator,.create_function=no_function_generator};
	(this->items)["withFloat"] = tl2::meta::tl_item{.tag=0x071b8685,.annotations=0x0,.name="withFloat",.create_object=no_object_generator,.create_function=no_function_generator};
}
