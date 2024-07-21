#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../__common/types/dictionaryAny.hpp"


namespace tl2 { namespace cases_bytes { 
struct TestDictAny {
	::tl2::DictionaryAny<double, int32_t> dict{};

	std::string_view tl_name() const { return "cases_bytes.testDictAny"; }
	uint32_t tl_tag() const { return 0x5a5fce57; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases_bytes

