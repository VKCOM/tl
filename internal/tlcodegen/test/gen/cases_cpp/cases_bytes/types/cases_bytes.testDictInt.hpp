#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../__common/types/dictionaryFieldAny.hpp"


namespace tl2 { namespace cases_bytes { 
struct TestDictInt {
	std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>> dict;

	std::string_view tl_name() const { return "cases_bytes.testDictInt"; }
	uint32_t tl_tag() const { return 0x453ace07; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases_bytes

