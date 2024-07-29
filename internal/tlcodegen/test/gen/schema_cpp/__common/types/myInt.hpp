#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
struct MyInt {
	int32_t val1 = 0;

	std::string_view tl_name() const { return "myInt"; }
	uint32_t tl_tag() const { return 0xc12375b7; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

} // namespace tl2

