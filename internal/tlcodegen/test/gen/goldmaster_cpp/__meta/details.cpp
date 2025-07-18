// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#include "basictl/io_streams.h"

#include <map>

#include "__meta/headers.h"

namespace {
  struct tl_items {
    public:
      std::map<std::string, std::shared_ptr<::tlgen::meta::tl_item>> items;
      std::map<uint32_t, std::shared_ptr<::tlgen::meta::tl_item>> items_by_tag;
      tl_items();
  };
    
  tl_items items;
    std::function<std::unique_ptr<::tlgen::meta::tl_object>()> no_object_generator = []() -> std::unique_ptr<::tlgen::meta::tl_object> {
        throw std::runtime_error("no generation for this type of objects");
    };
    std::function<std::unique_ptr<::tlgen::meta::tl_function>()> no_function_generator = []() -> std::unique_ptr<::tlgen::meta::tl_function> {
        throw std::runtime_error("no generation for this type of functions");
    };
}

std::optional<::tlgen::meta::tl_item> tlgen::meta::get_item_by_name(std::string &&s) {
  auto item = items.items.find(s);
  if (item != items.items.end()) {
    return *item->second;
  }
  return {};
}

std::optional<::tlgen::meta::tl_item> tlgen::meta::get_item_by_tag(std::uint32_t &&tag) {
  auto item = items.items_by_tag.find(tag);
  if (item != items.items_by_tag.end()) {
    return *item->second;
  }
  return {};
}

void tlgen::meta::set_create_object_by_name(std::string &&s, std::function<std::unique_ptr<::tlgen::meta::tl_object>()>&& gen) {
  auto item = items.items.find(s);
  if (item != items.items.end()) {
    item->second->has_create_object = true;
    item->second->create_object = gen;
    return;  
  }
  throw std::runtime_error("no item with such name + \"" + s + "\"");
}

void tlgen::meta::set_create_function_by_name(std::string &&s, std::function<std::unique_ptr<::tlgen::meta::tl_function>()>&& gen) {
  auto item = items.items.find(s);
  if (item != items.items.end()) {
    item->second->has_create_function = true;
    item->second->create_function = gen;
    return;  
  }
  throw std::runtime_error("no item with such name + \"" + s + "\"");
}

tl_items::tl_items() {
  auto item1647534323 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x623360f3,.annotations=0x0,.name="a.blue",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["a.blue"] = item1647534323;
  (this->items_by_tag)[0x623360f3] = item1647534323;
  auto item4082989673 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xf35d7a69,.annotations=0x0,.name="a.color",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["a.color"] = item4082989673;
  (this->items_by_tag)[0xf35d7a69] = item4082989673;
  auto item1630005176 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x6127e7b8,.annotations=0x0,.name="a.green",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["a.green"] = item1630005176;
  (this->items_by_tag)[0x6127e7b8] = item1630005176;
  auto item3090838077 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xb83a723d,.annotations=0x0,.name="a.red",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["a.red"] = item3090838077;
  (this->items_by_tag)[0xb83a723d] = item3090838077;
  auto item1887621519 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x7082d18f,.annotations=0x0,.name="a.top2",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["a.top2"] = item1887621519;
  (this->items_by_tag)[0x7082d18f] = item1887621519;
  auto item2808490051 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xa7662843,.annotations=0x0,.name="a.uNionA",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["a.uNionA"] = item2808490051;
  (this->items_by_tag)[0xa7662843] = item2808490051;
  auto item2487921303 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x944aaa97,.annotations=0x0,.name="ab.alias",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.alias"] = item2487921303;
  (this->items_by_tag)[0x944aaa97] = item2487921303;
  auto item549845805 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x20c5fb2d,.annotations=0x1,.name="ab.call1",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.call1"] = item549845805;
  (this->items_by_tag)[0x20c5fb2d] = item549845805;
  auto item2377295096 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x8db2a4f8,.annotations=0x1,.name="ab.call10",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.call10"] = item2377295096;
  (this->items_by_tag)[0x8db2a4f8] = item2377295096;
  auto item3971130220 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xecb2a36c,.annotations=0x1,.name="ab.call11",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.call11"] = item3971130220;
  (this->items_by_tag)[0xecb2a36c] = item3971130220;
  auto item2010509399 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x77d5f057,.annotations=0x1,.name="ab.call2",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.call2"] = item2010509399;
  (this->items_by_tag)[0x77d5f057] = item2010509399;
  auto item168309829 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x0a083445,.annotations=0x1,.name="ab.call3",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.call3"] = item168309829;
  (this->items_by_tag)[0x0a083445] = item168309829;
  auto item3240233502 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xc1220a1e,.annotations=0x1,.name="ab.call4",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.call4"] = item3240233502;
  (this->items_by_tag)[0xc1220a1e] = item3240233502;
  auto item2074399373 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x7ba4d28d,.annotations=0x1,.name="ab.call5",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.call5"] = item2074399373;
  (this->items_by_tag)[0x7ba4d28d] = item2074399373;
  auto item2228753867 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x84d815cb,.annotations=0x1,.name="ab.call6",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.call6"] = item2228753867;
  (this->items_by_tag)[0x84d815cb] = item2228753867;
  auto item1189875903 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x46ec10bf,.annotations=0x1,.name="ab.call7",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.call7"] = item1189875903;
  (this->items_by_tag)[0x46ec10bf] = item1189875903;
  auto item461787865 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x1b8652d9,.annotations=0x1,.name="ab.call8",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.call8"] = item461787865;
  (this->items_by_tag)[0x1b8652d9] = item461787865;
  auto item1977520236 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x75de906c,.annotations=0x1,.name="ab.call9",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.call9"] = item1977520236;
  (this->items_by_tag)[0x75de906c] = item1977520236;
  auto item1985065388 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x7651b1ac,.annotations=0x0,.name="ab.code",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.code"] = item1985065388;
  (this->items_by_tag)[0x7651b1ac] = item1985065388;
  auto item346250624 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x14a35d80,.annotations=0x0,.name="ab.counterChangeRequestPeriodsMany",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.counterChangeRequestPeriodsMany"] = item346250624;
  (this->items_by_tag)[0x14a35d80] = item346250624;
  auto item3653463525 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xd9c36de5,.annotations=0x0,.name="ab.counterChangeRequestPeriodsOne",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.counterChangeRequestPeriodsOne"] = item3653463525;
  (this->items_by_tag)[0xd9c36de5] = item3653463525;
  auto item516335166 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x1ec6a63e,.annotations=0x0,.name="ab.empty",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.empty"] = item516335166;
  (this->items_by_tag)[0x1ec6a63e] = item516335166;
  auto item3773394054 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xe0e96c86,.annotations=0x0,.name="ab.myType",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.myType"] = item3773394054;
  (this->items_by_tag)[0xe0e96c86] = item3773394054;
  auto item1303136554 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x4dac492a,.annotations=0x0,.name="ab.testMaybe",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.testMaybe"] = item1303136554;
  (this->items_by_tag)[0x4dac492a] = item1303136554;
  auto item3866873384 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xe67bce28,.annotations=0x0,.name="ab.topLevel1",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.topLevel1"] = item3866873384;
  (this->items_by_tag)[0xe67bce28] = item3866873384;
  auto item3472438267 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xcef933fb,.annotations=0x0,.name="ab.topLevel2",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.topLevel2"] = item3472438267;
  (this->items_by_tag)[0xcef933fb] = item3472438267;
  auto item2845831018 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xa99fef6a,.annotations=0x0,.name="ab.typeA",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.typeA"] = item2845831018;
  (this->items_by_tag)[0xa99fef6a] = item2845831018;
  auto item4281232728 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xff2e6d58,.annotations=0x0,.name="ab.typeB",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.typeB"] = item4281232728;
  (this->items_by_tag)[0xff2e6d58] = item4281232728;
  auto item1771179374 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x69920d6e,.annotations=0x0,.name="ab.typeC",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.typeC"] = item1771179374;
  (this->items_by_tag)[0x69920d6e] = item1771179374;
  auto item1986092017 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x76615bf1,.annotations=0x0,.name="ab.typeD",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.typeD"] = item1986092017;
  (this->items_by_tag)[0x76615bf1] = item1986092017;
  auto item1902670721 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x71687381,.annotations=0x0,.name="ab.useCycle",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.useCycle"] = item1902670721;
  (this->items_by_tag)[0x71687381] = item1902670721;
  auto item858118276 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x3325d884,.annotations=0x0,.name="ab.useDictString",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["ab.useDictString"] = item858118276;
  (this->items_by_tag)[0x3325d884] = item858118276;
  auto item3747739186 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xdf61f632,.annotations=0x0,.name="au.nionA",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["au.nionA"] = item3747739186;
  (this->items_by_tag)[0xdf61f632] = item3747739186;
  auto item2840008772 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xa9471844,.annotations=0x0,.name="b.red",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["b.red"] = item2840008772;
  (this->items_by_tag)[0xa9471844] = item2840008772;
  auto item3937838772 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xeab6a6b4,.annotations=0x0,.name="cd.myType",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["cd.myType"] = item3937838772;
  (this->items_by_tag)[0xeab6a6b4] = item3937838772;
  auto item2350919524 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x8c202f64,.annotations=0x0,.name="cd.response",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["cd.response"] = item2350919524;
  (this->items_by_tag)[0x8c202f64] = item2350919524;
  auto item1557252745 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x5cd1ca89,.annotations=0x0,.name="cd.topLevel3",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["cd.topLevel3"] = item1557252745;
  (this->items_by_tag)[0x5cd1ca89] = item1557252745;
  auto item2821826848 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xa831a920,.annotations=0x0,.name="cd.typeA",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["cd.typeA"] = item2821826848;
  (this->items_by_tag)[0xa831a920] = item2821826848;
  auto item930826646 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x377b4996,.annotations=0x0,.name="cd.typeB",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["cd.typeB"] = item930826646;
  (this->items_by_tag)[0x377b4996] = item930826646;
  auto item3675231188 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xdb0f93d4,.annotations=0x0,.name="cd.typeC",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["cd.typeC"] = item3675231188;
  (this->items_by_tag)[0xdb0f93d4] = item3675231188;
  auto item3042083461 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xb5528285,.annotations=0x0,.name="cd.typeD",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["cd.typeD"] = item3042083461;
  (this->items_by_tag)[0xb5528285] = item3042083461;
  auto item1859550368 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x6ed67ca0,.annotations=0x0,.name="cd.useCycle",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["cd.useCycle"] = item1859550368;
  (this->items_by_tag)[0x6ed67ca0] = item1859550368;
  auto item326028446 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x136ecc9e,.annotations=0x0,.name="cyc1.myCycle",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["cyc1.myCycle"] = item326028446;
  (this->items_by_tag)[0x136ecc9e] = item326028446;
  auto item4221955787 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xfba5eecb,.annotations=0x0,.name="cyc2.myCycle",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["cyc2.myCycle"] = item4221955787;
  (this->items_by_tag)[0xfba5eecb] = item4221955787;
  auto item1199990880 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x47866860,.annotations=0x0,.name="cyc3.myCycle",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["cyc3.myCycle"] = item1199990880;
  (this->items_by_tag)[0x47866860] = item1199990880;
  auto item3362257635 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xc867fae3,.annotations=0x0,.name="cycleTuple",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["cycleTuple"] = item3362257635;
  (this->items_by_tag)[0xc867fae3] = item3362257635;
  auto item1685969653 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x647ddaf5,.annotations=0x0,.name="halfStr",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["halfStr"] = item1685969653;
  (this->items_by_tag)[0x647ddaf5] = item1685969653;
  auto item313217561 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x12ab5219,.annotations=0x0,.name="hren",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["hren"] = item313217561;
  (this->items_by_tag)[0x12ab5219] = item313217561;
  auto item2823855066 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xa8509bda,.annotations=0x0,.name="int",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["int"] = item2823855066;
  (this->items_by_tag)[0xa8509bda] = item2823855066;
  auto item2033510175 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x7934e71f,.annotations=0x0,.name="int32",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["int32"] = item2033510175;
  (this->items_by_tag)[0x7934e71f] = item2033510175;
  auto item4116749792 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xf5609de0,.annotations=0x0,.name="int64",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["int64"] = item4116749792;
  (this->items_by_tag)[0xf5609de0] = item4116749792;
  auto item570911930 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x22076cba,.annotations=0x0,.name="long",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["long"] = item570911930;
  (this->items_by_tag)[0x22076cba] = item570911930;
  auto item3294066236 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xc457763c,.annotations=0x0,.name="maybeTest1",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["maybeTest1"] = item3294066236;
  (this->items_by_tag)[0xc457763c] = item3294066236;
  auto item236644382 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x0e1ae81e,.annotations=0x0,.name="multiPoint",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["multiPoint"] = item236644382;
  (this->items_by_tag)[0x0e1ae81e] = item236644382;
  auto item3126452561 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xba59e151,.annotations=0x0,.name="myInt32",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["myInt32"] = item3126452561;
  (this->items_by_tag)[0xba59e151] = item3126452561;
  auto item496360349 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x1d95db9d,.annotations=0x0,.name="myInt64",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["myInt64"] = item496360349;
  (this->items_by_tag)[0x1d95db9d] = item496360349;
  auto item3322682177 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xc60c1b41,.annotations=0x0,.name="myNat",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["myNat"] = item3322682177;
  (this->items_by_tag)[0xc60c1b41] = item3322682177;
  auto item2044774111 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x79e0c6df,.annotations=0x0,.name="myPlus",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["myPlus"] = item2044774111;
  (this->items_by_tag)[0x79e0c6df] = item2044774111;
  auto item1764501787 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x692c291b,.annotations=0x0,.name="myPlus3",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["myPlus3"] = item1764501787;
  (this->items_by_tag)[0x692c291b] = item1764501787;
  auto item2374402937 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x8d868379,.annotations=0x0,.name="myZero",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["myZero"] = item2374402937;
  (this->items_by_tag)[0x8d868379] = item2374402937;
  auto item272253135 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x103a40cf,.annotations=0x0,.name="myZero3",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["myZero3"] = item272253135;
  (this->items_by_tag)[0x103a40cf] = item272253135;
  auto item877518672 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x344ddf50,.annotations=0x0,.name="nativeWrappers",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["nativeWrappers"] = item877518672;
  (this->items_by_tag)[0x344ddf50] = item877518672;
  auto item980583204 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x3a728324,.annotations=0x0,.name="noStr",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["noStr"] = item980583204;
  (this->items_by_tag)[0x3a728324] = item980583204;
  auto item842905150 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x323db63e,.annotations=0x0,.name="replace",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["replace"] = item842905150;
  (this->items_by_tag)[0x323db63e] = item842905150;
  auto item4236374024 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xfc81f008,.annotations=0x0,.name="replace10",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["replace10"] = item4236374024;
  (this->items_by_tag)[0xfc81f008] = item4236374024;
  auto item3960606868 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xec121094,.annotations=0x0,.name="replace12",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["replace12"] = item3960606868;
  (this->items_by_tag)[0xec121094] = item3960606868;
  auto item578872368 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x2280e430,.annotations=0x0,.name="replace15",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["replace15"] = item578872368;
  (this->items_by_tag)[0x2280e430] = item578872368;
  auto item4100955035 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xf46f9b9b,.annotations=0x0,.name="replace17",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["replace17"] = item4100955035;
  (this->items_by_tag)[0xf46f9b9b] = item4100955035;
  auto item1884149522 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x704dd712,.annotations=0x0,.name="replace18",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["replace18"] = item1884149522;
  (this->items_by_tag)[0x704dd712] = item1884149522;
  auto item3805604846 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xe2d4ebee,.annotations=0x0,.name="replace2",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["replace2"] = item3805604846;
  (this->items_by_tag)[0xe2d4ebee] = item3805604846;
  auto item1373840612 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x51e324e4,.annotations=0x0,.name="replace3",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["replace3"] = item1373840612;
  (this->items_by_tag)[0x51e324e4] = item1373840612;
  auto item2338047882 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x8b5bc78a,.annotations=0x0,.name="replace5",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["replace5"] = item2338047882;
  (this->items_by_tag)[0x8b5bc78a] = item2338047882;
  auto item2882837766 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xabd49d06,.annotations=0x0,.name="replace6",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["replace6"] = item2882837766;
  (this->items_by_tag)[0xabd49d06] = item2882837766;
  auto item4106644895 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xf4c66d9f,.annotations=0x0,.name="replace7",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["replace7"] = item4106644895;
  (this->items_by_tag)[0xf4c66d9f] = item4106644895;
  auto item3592864023 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xd626c117,.annotations=0x0,.name="replace8",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["replace8"] = item3592864023;
  (this->items_by_tag)[0xd626c117] = item3592864023;
  auto item2513803461 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x95d598c5,.annotations=0x0,.name="replace9",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["replace9"] = item2513803461;
  (this->items_by_tag)[0x95d598c5] = item2513803461;
  auto item4287593912 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xff8f7db8,.annotations=0x0,.name="service5.emptyOutput",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["service5.emptyOutput"] = item4287593912;
  (this->items_by_tag)[0xff8f7db8] = item4287593912;
  auto item4287593913 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xff8f7db9,.annotations=0x0,.name="service5Long.emptyOutput",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["service5Long.emptyOutput"] = item4287593913;
  (this->items_by_tag)[0xff8f7db9] = item4287593913;
  auto item3692498933 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xdc170ff5,.annotations=0x0,.name="service5Long.stringOutput",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["service5Long.stringOutput"] = item3692498933;
  (this->items_by_tag)[0xdc170ff5] = item3692498933;
  auto item3692498932 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xdc170ff4,.annotations=0x0,.name="service5.stringOutput",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["service5.stringOutput"] = item3692498932;
  (this->items_by_tag)[0xdc170ff4] = item3692498932;
  auto item3039325732 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xb5286e24,.annotations=0x0,.name="string",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["string"] = item3039325732;
  (this->items_by_tag)[0xb5286e24] = item3039325732;
  auto item2291273360 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x88920e90,.annotations=0x0,.name="testMaybe",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["testMaybe"] = item2291273360;
  (this->items_by_tag)[0x88920e90] = item2291273360;
  auto item178273522 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x0aa03cf2,.annotations=0x0,.name="testMaybe2",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["testMaybe2"] = item178273522;
  (this->items_by_tag)[0x0aa03cf2] = item178273522;
  auto item1072550713 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x3fedd339,.annotations=0x0,.name="true",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["true"] = item1072550713;
  (this->items_by_tag)[0x3fedd339] = item1072550713;
  auto item360084417 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x157673c1,.annotations=0x0,.name="typeA",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["typeA"] = item360084417;
  (this->items_by_tag)[0x157673c1] = item360084417;
  auto item2634172418 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x9d024802,.annotations=0x0,.name="typeB",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["typeB"] = item2634172418;
  (this->items_by_tag)[0x9d024802] = item2634172418;
  auto item1804530751 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x6b8ef43f,.annotations=0x0,.name="typeC",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["typeC"] = item1804530751;
  (this->items_by_tag)[0x6b8ef43f] = item1804530751;
  auto item2985571998 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xb1f4369e,.annotations=0x0,.name="typeD",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["typeD"] = item2985571998;
  (this->items_by_tag)[0xb1f4369e] = item2985571998;
  auto item1948344786 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x742161d2,.annotations=0x0,.name="unionArgsUse",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["unionArgsUse"] = item1948344786;
  (this->items_by_tag)[0x742161d2] = item1948344786;
  auto item4221364247 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xfb9ce817,.annotations=0x0,.name="useDictUgly",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["useDictUgly"] = item4221364247;
  (this->items_by_tag)[0xfb9ce817] = item4221364247;
  auto item174320735 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x0a63ec5f,.annotations=0x0,.name="useResponse",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["useResponse"] = item174320735;
  (this->items_by_tag)[0x0a63ec5f] = item174320735;
  auto item2594430693 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x9aa3dee5,.annotations=0x0,.name="useStr",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["useStr"] = item2594430693;
  (this->items_by_tag)[0x9aa3dee5] = item2594430693;
  auto item3755819392 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0xdfdd4180,.annotations=0x0,.name="useTrue",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["useTrue"] = item3755819392;
  (this->items_by_tag)[0xdfdd4180] = item3755819392;
  auto item1015381586 = std::shared_ptr<::tlgen::meta::tl_item>(new ::tlgen::meta::tl_item{.tag=0x3c857e52,.annotations=0x2,.name="usefulService.getUserEntity",.create_object=no_object_generator,.create_function=no_function_generator});
  (this->items)["usefulService.getUserEntity"] = item1015381586;
  (this->items_by_tag)[0x3c857e52] = item1015381586;
}
