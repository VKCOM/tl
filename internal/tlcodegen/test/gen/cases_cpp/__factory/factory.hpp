#pragma once
#include "../__meta/meta.hpp"

#include "../cases/types/cases.testVector.hpp"
#include "../cases/types/cases.testUnionContainer.hpp"
#include "../cases/types/cases.testUnion2.hpp"
#include "../cases/types/cases.testUnion1.hpp"
#include "../cases/types/cases.testTuple.hpp"
#include "../cases/types/cases.testRecursiveFieldMask.hpp"
#include "../cases/types/cases.testOutFieldMaskContainer.hpp"
#include "../cases/types/cases.testMaybe.hpp"
#include "../cases/types/cases.testLocalFieldmask.hpp"
#include "../cases/types/cases.testEnumContainer.hpp"
#include "../cases/types/cases.testDictString.hpp"
#include "../cases/types/cases.testDictInt.hpp"
#include "../cases/types/cases.testDictAny.hpp"
#include "../cases/types/cases.testBeforeReadBitValidation.hpp"
#include "../cases/types/cases.testArray.hpp"
#include "../cases/types/cases.testAllPossibleFieldConfigsContainer.hpp"
#include "../cases/types/cases.replace7plusplus.hpp"
#include "../cases/types/cases.replace7plus.hpp"
#include "../cases/types/cases.replace7.hpp"
#include "../cases/types/cases.myCycle1.hpp"
#include "../cases/types/cases.myCycle2.hpp"
#include "../cases/types/cases.myCycle3.hpp"
#include "../cases_bytes/types/cases_bytes.testVector.hpp"
#include "../cases_bytes/types/cases_bytes.testTuple.hpp"
#include "../cases_bytes/types/cases_bytes.testEnumContainer.hpp"
#include "../cases/types/cases.TestEnumItems.hpp"
#include "../cases_bytes/types/cases_bytes.TestEnumItems.hpp"
#include "../cases_bytes/types/cases_bytes.testDictStringString.hpp"
#include "../cases_bytes/types/cases_bytes.testDictString.hpp"
#include "../cases_bytes/types/cases_bytes.testDictInt.hpp"
#include "../cases_bytes/types/cases_bytes.testDictAny.hpp"
#include "../cases_bytes/types/cases_bytes.testArray.hpp"
#include "../benchmarks/types/benchmarks.vrutoyTopLevelContainerWithDependency.hpp"
#include "../benchmarks/types/benchmarks.vrutoyTopLevelContainer.hpp"
#include "../benchmarks/types/benchmarks.vrutoytopLevelUnionEmpty.hpp"
#include "../benchmarks/types/benchmarks.vrutoytopLevelUnionBig.hpp"
#include "../benchmarks/types/benchmarks.vruposition.hpp"
#include "../__common/types/true.hpp"
#include "../benchmarks/types/benchmarks.vruhash.hpp"


namespace tl2 {
namespace factory {
    void init_tl_create_objects() {
		tl2::meta::set_create_object_by_name("benchmarks.vruhash",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::benchmarks::Vruhash>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("benchmarks.vruposition",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::benchmarks::Vruposition>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("benchmarks.vrutoyTopLevelContainer",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::benchmarks::VrutoyTopLevelContainer>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("benchmarks.vrutoyTopLevelContainerWithDependency",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::benchmarks::VrutoyTopLevelContainerWithDependency>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("benchmarks.vrutoytopLevelUnionBig",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::benchmarks::VrutoytopLevelUnionBig>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("benchmarks.vrutoytopLevelUnionEmpty",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::benchmarks::VrutoytopLevelUnionEmpty>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases_bytes.testArray",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases_bytes::TestArray>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases_bytes.testDictAny",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases_bytes::TestDictAny>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases_bytes.testDictInt",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases_bytes::TestDictInt>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases_bytes.testDictString",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases_bytes::TestDictString>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases_bytes.testDictStringString",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases_bytes::TestDictStringString>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases_bytes.testEnum1",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases_bytes::TestEnum1>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases_bytes.testEnum2",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases_bytes::TestEnum2>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases_bytes.testEnum3",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases_bytes::TestEnum3>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases_bytes.testEnumContainer",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases_bytes::TestEnumContainer>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases_bytes.testTuple",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases_bytes::TestTuple>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases_bytes.testVector",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases_bytes::TestVector>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.myCycle1",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::MyCycle1>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.myCycle2",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::MyCycle2>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.myCycle3",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::MyCycle3>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.replace7",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::Replace7>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.replace7plus",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::Replace7plus>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.replace7plusplus",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::Replace7plusplus>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.testAllPossibleFieldConfigsContainer",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::TestAllPossibleFieldConfigsContainer>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.testArray",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::TestArray>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.testBeforeReadBitValidation",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::TestBeforeReadBitValidation>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.testDictAny",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::TestDictAny>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.testDictInt",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::TestDictInt>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.testDictString",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::TestDictString>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.testEnum1",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::TestEnum1>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.testEnum2",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::TestEnum2>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.testEnum3",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::TestEnum3>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.testEnumContainer",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::TestEnumContainer>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.testLocalFieldmask",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::TestLocalFieldmask>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.testMaybe",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::TestMaybe>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.testOutFieldMaskContainer",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::TestOutFieldMaskContainer>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.testRecursiveFieldMask",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::TestRecursiveFieldMask>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.testTuple",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::TestTuple>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.testUnion1",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::TestUnion1>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.testUnion2",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::TestUnion2>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.testUnionContainer",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::TestUnionContainer>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("cases.testVector",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::cases::TestVector>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
		tl2::meta::set_create_object_by_name("true",[]() -> tl2::meta::tl_object {
        auto obj = std::make_shared<tl2::True>();
        return tl2::meta::tl_object{
                .read=[obj](auto &in) -> bool { return obj->read(in); },
                .write=[obj](auto &out) -> bool { return obj->write(out); },
                .read_boxed=[obj](auto &in) -> bool { return obj->read_boxed(in); },
                .write_boxed=[obj](auto &out) -> bool { return obj->write_boxed(out); },
        };
    });
	}
};
};
