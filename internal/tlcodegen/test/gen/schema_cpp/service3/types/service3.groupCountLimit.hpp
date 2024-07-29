#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace service3 { 
struct GroupCountLimit {
	std::vector<int32_t> types;
	int32_t limit = 0;

	std::string_view tl_name() const { return "service3.groupCountLimit"; }
	uint32_t tl_tag() const { return 0x8c04ea7f; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::service3

