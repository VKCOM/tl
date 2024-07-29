#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
struct BoxedString {
	std::string x;

	std::string_view tl_name() const { return "boxedString"; }
	uint32_t tl_tag() const { return 0x548994db; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::string & result);
	bool write_result(::basictl::tl_ostream & s, std::string & result);
};

} // namespace tl2

