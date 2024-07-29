#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
struct NonOptNat {
	uint32_t n = 0;
	std::vector<int32_t> xs;

	std::string_view tl_name() const { return "nonOptNat"; }
	uint32_t tl_tag() const { return 0x45366605; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

} // namespace tl2

