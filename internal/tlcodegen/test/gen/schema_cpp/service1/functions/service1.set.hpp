#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace service1 { 
struct Set {
	std::string key;
	int32_t flags = 0;
	int32_t delay = 0;
	std::string value;

	std::string_view tl_name() const { return "service1.set"; }
	uint32_t tl_tag() const { return 0x05ae5f66; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, bool & result);
	bool write_result(::basictl::tl_ostream & s, bool & result);
};

}} // namespace tl2::service1

