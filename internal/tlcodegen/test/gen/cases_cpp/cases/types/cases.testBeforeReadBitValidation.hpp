#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace cases { 
struct TestBeforeReadBitValidation {
	uint32_t n = 0;
	std::vector<int32_t> a;
	std::vector<int32_t> b;

	std::string_view tl_name() const { return "cases.testBeforeReadBitValidation"; }
	uint32_t tl_tag() const { return 0x9b2396db; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases

