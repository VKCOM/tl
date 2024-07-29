#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/myBoxedTupleSlice.hpp"


namespace tl2 { 
struct BoxedTupleSlice2 {
	::tl2::MyBoxedTupleSlice x{};

	std::string_view tl_name() const { return "boxedTupleSlice2"; }
	uint32_t tl_tag() const { return 0x1cdf4705; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::MyBoxedTupleSlice & result);
	bool write_result(::basictl::tl_ostream & s, ::tl2::MyBoxedTupleSlice & result);
};

} // namespace tl2

