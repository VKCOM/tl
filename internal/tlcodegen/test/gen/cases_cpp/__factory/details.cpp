#include "../__meta/headers.hpp"
#include "headers.hpp"

#include "../cases/types/cases.testVector.hpp"
#include "../cases/types/cases.testUnionContainer.hpp"
#include "../cases/types/cases.testUnion2.hpp"
#include "../cases/types/cases.testUnion1.hpp"
#include "../cases/types/cases.testTuple.hpp"
#include "../cases/types/cases.testRecursiveFieldMask.hpp"
#include "../cases/types/cases.testOutFieldMaskContainer.hpp"
#include "../cases/types/cases.testMaybe.hpp"
#include "../cases/types/cases.testLocalFieldmask.hpp"
#include "../cases/types/cases.testInplaceStructArgs2.hpp"
#include "../cases/types/cases.testInplaceStructArgs.hpp"
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
#include "../__common_namespace/types/true.hpp"
#include "../benchmarks/types/benchmarks.vruhash.hpp"

void tl2::factory::set_all_factories() {

	struct tl2_benchmarks_Vruhash_tl_object : public tl2::meta::tl_object {
        tl2::benchmarks::Vruhash object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("benchmarks.vruhash", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_benchmarks_Vruhash_tl_object>();
	});

	struct tl2_benchmarks_Vruposition_tl_object : public tl2::meta::tl_object {
        tl2::benchmarks::Vruposition object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("benchmarks.vruposition", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_benchmarks_Vruposition_tl_object>();
	});

	struct tl2_benchmarks_VrutoyTopLevelContainer_tl_object : public tl2::meta::tl_object {
        tl2::benchmarks::VrutoyTopLevelContainer object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("benchmarks.vrutoyTopLevelContainer", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_benchmarks_VrutoyTopLevelContainer_tl_object>();
	});

	struct tl2_benchmarks_VrutoyTopLevelContainerWithDependency_tl_object : public tl2::meta::tl_object {
        tl2::benchmarks::VrutoyTopLevelContainerWithDependency object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("benchmarks.vrutoyTopLevelContainerWithDependency", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_benchmarks_VrutoyTopLevelContainerWithDependency_tl_object>();
	});

	struct tl2_benchmarks_VrutoytopLevelUnionBig_tl_object : public tl2::meta::tl_object {
        tl2::benchmarks::VrutoytopLevelUnionBig object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("benchmarks.vrutoytopLevelUnionBig", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_benchmarks_VrutoytopLevelUnionBig_tl_object>();
	});

	struct tl2_benchmarks_VrutoytopLevelUnionEmpty_tl_object : public tl2::meta::tl_object {
        tl2::benchmarks::VrutoytopLevelUnionEmpty object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("benchmarks.vrutoytopLevelUnionEmpty", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_benchmarks_VrutoytopLevelUnionEmpty_tl_object>();
	});

	struct tl2_cases_bytes_TestArray_tl_object : public tl2::meta::tl_object {
        tl2::cases_bytes::TestArray object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases_bytes.testArray", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_bytes_TestArray_tl_object>();
	});

	struct tl2_cases_bytes_TestDictAny_tl_object : public tl2::meta::tl_object {
        tl2::cases_bytes::TestDictAny object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases_bytes.testDictAny", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_bytes_TestDictAny_tl_object>();
	});

	struct tl2_cases_bytes_TestDictInt_tl_object : public tl2::meta::tl_object {
        tl2::cases_bytes::TestDictInt object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases_bytes.testDictInt", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_bytes_TestDictInt_tl_object>();
	});

	struct tl2_cases_bytes_TestDictString_tl_object : public tl2::meta::tl_object {
        tl2::cases_bytes::TestDictString object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases_bytes.testDictString", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_bytes_TestDictString_tl_object>();
	});

	struct tl2_cases_bytes_TestDictStringString_tl_object : public tl2::meta::tl_object {
        tl2::cases_bytes::TestDictStringString object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases_bytes.testDictStringString", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_bytes_TestDictStringString_tl_object>();
	});

	struct tl2_cases_bytes_TestEnum1_tl_object : public tl2::meta::tl_object {
        tl2::cases_bytes::TestEnum1 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases_bytes.testEnum1", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_bytes_TestEnum1_tl_object>();
	});

	struct tl2_cases_bytes_TestEnum2_tl_object : public tl2::meta::tl_object {
        tl2::cases_bytes::TestEnum2 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases_bytes.testEnum2", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_bytes_TestEnum2_tl_object>();
	});

	struct tl2_cases_bytes_TestEnum3_tl_object : public tl2::meta::tl_object {
        tl2::cases_bytes::TestEnum3 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases_bytes.testEnum3", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_bytes_TestEnum3_tl_object>();
	});

	struct tl2_cases_bytes_TestEnumContainer_tl_object : public tl2::meta::tl_object {
        tl2::cases_bytes::TestEnumContainer object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases_bytes.testEnumContainer", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_bytes_TestEnumContainer_tl_object>();
	});

	struct tl2_cases_bytes_TestTuple_tl_object : public tl2::meta::tl_object {
        tl2::cases_bytes::TestTuple object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases_bytes.testTuple", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_bytes_TestTuple_tl_object>();
	});

	struct tl2_cases_bytes_TestVector_tl_object : public tl2::meta::tl_object {
        tl2::cases_bytes::TestVector object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases_bytes.testVector", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_bytes_TestVector_tl_object>();
	});

	struct tl2_cases_MyCycle1_tl_object : public tl2::meta::tl_object {
        tl2::cases::MyCycle1 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.myCycle1", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_MyCycle1_tl_object>();
	});

	struct tl2_cases_MyCycle2_tl_object : public tl2::meta::tl_object {
        tl2::cases::MyCycle2 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.myCycle2", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_MyCycle2_tl_object>();
	});

	struct tl2_cases_MyCycle3_tl_object : public tl2::meta::tl_object {
        tl2::cases::MyCycle3 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.myCycle3", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_MyCycle3_tl_object>();
	});

	struct tl2_cases_Replace7_tl_object : public tl2::meta::tl_object {
        tl2::cases::Replace7 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.replace7", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_Replace7_tl_object>();
	});

	struct tl2_cases_Replace7plus_tl_object : public tl2::meta::tl_object {
        tl2::cases::Replace7plus object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.replace7plus", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_Replace7plus_tl_object>();
	});

	struct tl2_cases_Replace7plusplus_tl_object : public tl2::meta::tl_object {
        tl2::cases::Replace7plusplus object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.replace7plusplus", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_Replace7plusplus_tl_object>();
	});

	struct tl2_cases_TestAllPossibleFieldConfigsContainer_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestAllPossibleFieldConfigsContainer object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testAllPossibleFieldConfigsContainer", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestAllPossibleFieldConfigsContainer_tl_object>();
	});

	struct tl2_cases_TestArray_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestArray object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testArray", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestArray_tl_object>();
	});

	struct tl2_cases_TestBeforeReadBitValidation_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestBeforeReadBitValidation object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testBeforeReadBitValidation", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestBeforeReadBitValidation_tl_object>();
	});

	struct tl2_cases_TestDictAny_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestDictAny object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testDictAny", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestDictAny_tl_object>();
	});

	struct tl2_cases_TestDictInt_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestDictInt object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testDictInt", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestDictInt_tl_object>();
	});

	struct tl2_cases_TestDictString_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestDictString object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testDictString", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestDictString_tl_object>();
	});

	struct tl2_cases_TestEnum1_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestEnum1 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testEnum1", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestEnum1_tl_object>();
	});

	struct tl2_cases_TestEnum2_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestEnum2 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testEnum2", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestEnum2_tl_object>();
	});

	struct tl2_cases_TestEnum3_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestEnum3 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testEnum3", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestEnum3_tl_object>();
	});

	struct tl2_cases_TestEnumContainer_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestEnumContainer object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testEnumContainer", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestEnumContainer_tl_object>();
	});

	struct tl2_cases_TestInplaceStructArgs_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestInplaceStructArgs object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testInplaceStructArgs", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestInplaceStructArgs_tl_object>();
	});

	struct tl2_cases_TestInplaceStructArgs2_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestInplaceStructArgs2 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testInplaceStructArgs2", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestInplaceStructArgs2_tl_object>();
	});

	struct tl2_cases_TestLocalFieldmask_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestLocalFieldmask object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testLocalFieldmask", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestLocalFieldmask_tl_object>();
	});

	struct tl2_cases_TestMaybe_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestMaybe object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testMaybe", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestMaybe_tl_object>();
	});

	struct tl2_cases_TestOutFieldMaskContainer_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestOutFieldMaskContainer object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testOutFieldMaskContainer", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestOutFieldMaskContainer_tl_object>();
	});

	struct tl2_cases_TestRecursiveFieldMask_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestRecursiveFieldMask object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testRecursiveFieldMask", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestRecursiveFieldMask_tl_object>();
	});

	struct tl2_cases_TestTuple_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestTuple object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testTuple", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestTuple_tl_object>();
	});

	struct tl2_cases_TestUnion1_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestUnion1 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testUnion1", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestUnion1_tl_object>();
	});

	struct tl2_cases_TestUnion2_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestUnion2 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testUnion2", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestUnion2_tl_object>();
	});

	struct tl2_cases_TestUnionContainer_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestUnionContainer object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testUnionContainer", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestUnionContainer_tl_object>();
	});

	struct tl2_cases_TestVector_tl_object : public tl2::meta::tl_object {
        tl2::cases::TestVector object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("cases.testVector", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_cases_TestVector_tl_object>();
	});

	struct tl2_True_tl_object : public tl2::meta::tl_object {
        tl2::True object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("true", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_True_tl_object>();
	});

}
