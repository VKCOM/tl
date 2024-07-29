#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
struct BoxedVector32BoxedElem {
	std::vector<int32_t> x;

	std::string_view tl_name() const { return "boxedVector32BoxedElem"; }
	uint32_t tl_tag() const { return 0x591cecd4; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::vector<int32_t> & result);
	bool write_result(::basictl::tl_ostream & s, std::vector<int32_t> & result);
};

} // namespace tl2

