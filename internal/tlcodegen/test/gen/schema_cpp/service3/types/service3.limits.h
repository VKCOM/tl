#pragma once

#include "../../basictl/io_streams.h"
#include "service3.groupSizeLimit.h"
#include "service3.groupCountLimit.h"


namespace tl2 { namespace service3 { 
struct Limits {
	int32_t default_group_size_limit = 0;
	std::vector<::tl2::service3::GroupSizeLimit> custom_group_size_limits;
	int32_t default_group_count_limit = 0;
	std::vector<::tl2::service3::GroupCountLimit> custom_group_count_limits;

	std::string_view tl_name() const { return "service3.limits"; }
	uint32_t tl_tag() const { return 0x80ee61ca; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const Limits& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service3

