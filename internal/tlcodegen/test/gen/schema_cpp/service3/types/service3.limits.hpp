#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "service3.groupSizeLimit.hpp"
#include "service3.groupCountLimit.hpp"


namespace tl2 { namespace service3 { 
struct Limits {
	int32_t default_group_size_limit = 0;
	std::vector<::tl2::service3::GroupSizeLimit> custom_group_size_limits;
	int32_t default_group_count_limit = 0;
	std::vector<::tl2::service3::GroupCountLimit> custom_group_count_limits;

	std::string_view tl_name() const { return "service3.limits"; }
	uint32_t tl_tag() const { return 0x80ee61ca; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::service3

