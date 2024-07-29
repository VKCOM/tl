#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace service1 { 
struct GetWildcardList {
	std::string prefix;

	std::string_view tl_name() const { return "service1.getWildcardList"; }
	uint32_t tl_tag() const { return 0x56b6ead4; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::vector<std::string> & result);
	bool write_result(::basictl::tl_ostream & s, std::vector<std::string> & result);
};

}} // namespace tl2::service1

