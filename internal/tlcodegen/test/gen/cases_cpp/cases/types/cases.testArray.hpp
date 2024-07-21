#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace cases { 
struct TestArray {
	uint32_t n = 0;
	std::vector<int32_t> arr;

	std::string_view tl_name() const { return "cases.testArray"; }
	uint32_t tl_tag() const { return 0xa888030d; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases

