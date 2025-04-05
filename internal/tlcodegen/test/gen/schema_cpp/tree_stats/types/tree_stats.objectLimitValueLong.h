#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { namespace tree_stats { 
struct ObjectLimitValueLong {

	std::string_view tl_name() const { return "tree_stats.objectLimitValueLong"; }
	uint32_t tl_tag() const { return 0x73111993; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const ObjectLimitValueLong& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::tree_stats

