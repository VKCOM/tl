#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../__common_namespace/types/dictionary.hpp"


namespace tl2 { namespace cases_bytes { 
struct TestDictString {
	::tl2::Dictionary<int32_t> dict{};

	std::string_view tl_name() const { return "cases_bytes.testDictString"; }
	uint32_t tl_tag() const { return 0x6c04d6ce; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases_bytes

