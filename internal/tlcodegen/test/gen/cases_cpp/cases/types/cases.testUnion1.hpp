#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace cases { 
struct TestUnion1 {
	int32_t value = 0;

	std::string_view tl_name() const { return "cases.testUnion1"; }
	uint32_t tl_tag() const { return 0x4b4f09b1; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases

