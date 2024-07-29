#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace service1 { 
struct Exists {
	std::string key;

	std::string_view tl_name() const { return "service1.exists"; }
	uint32_t tl_tag() const { return 0xe0284c9e; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, bool & result);
	bool write_result(::basictl::tl_ostream & s, bool & result);
};

}} // namespace tl2::service1

