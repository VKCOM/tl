#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace cases_bytes { 
struct TestVector {
	std::vector<std::string> arr;

	std::string_view tl_name() const { return "cases_bytes.testVector"; }
	uint32_t tl_tag() const { return 0x3647c8ae; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases_bytes

