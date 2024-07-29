#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
struct GetNonOptNat {
	uint32_t n = 0;
	std::vector<int32_t> xs;

	std::string_view tl_name() const { return "getNonOptNat"; }
	uint32_t tl_tag() const { return 0x67665961; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::vector<int32_t> & result);
	bool write_result(::basictl::tl_ostream & s, std::vector<int32_t> & result);
};

} // namespace tl2

