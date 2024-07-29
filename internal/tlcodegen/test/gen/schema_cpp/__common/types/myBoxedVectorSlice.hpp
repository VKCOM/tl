#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
struct MyBoxedVectorSlice {
	std::vector<int32_t> data;

	std::string_view tl_name() const { return "myBoxedVectorSlice"; }
	uint32_t tl_tag() const { return 0x57d164bb; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

} // namespace tl2

