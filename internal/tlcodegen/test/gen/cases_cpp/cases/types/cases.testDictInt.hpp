#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../__common_namespace/types/dictionaryFieldAny.hpp"


namespace tl2 { namespace cases { 
struct TestDictInt {
	std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>> dict;

	std::string_view tl_name() const { return "cases.testDictInt"; }
	uint32_t tl_tag() const { return 0xd3877643; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases

