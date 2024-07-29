#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
struct MyBoxedTupleSlice {
	uint32_t n = 0;
	std::vector<int32_t> data;

	std::string_view tl_name() const { return "myBoxedTupleSlice"; }
	uint32_t tl_tag() const { return 0x25d1a1be; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

} // namespace tl2

