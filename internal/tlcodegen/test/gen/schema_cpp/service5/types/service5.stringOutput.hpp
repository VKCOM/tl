#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace service5 { 
struct StringOutput {
	int32_t http_code = 0;
	std::string response;

	std::string_view tl_name() const { return "service5.stringOutput"; }
	uint32_t tl_tag() const { return 0x179e9863; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::service5

