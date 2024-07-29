#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace service6 { 
struct Error {
	int32_t code = 0;

	std::string_view tl_name() const { return "service6.error"; }
	uint32_t tl_tag() const { return 0x738553ef; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::service6

