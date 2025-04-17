#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "tree_stats/types/tree_stats.objectLimitValueLong.h"
#include "tree_stats/types/tree_stats.objectLimitValueDouble.h"


namespace tl2 { namespace tree_stats { 
struct ObjectLimitValue {
	std::variant<::tl2::tree_stats::ObjectLimitValueLong, ::tl2::tree_stats::ObjectLimitValueDouble> value;

	bool is_Long() const { return value.index() == 0; }
	bool is_Double() const { return value.index() == 1; }

	void set_Long() { value.emplace<0>(); }

	std::string_view tl_name() const;
	uint32_t tl_tag() const;

	bool write_json(std::ostream& s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;
};

}} // namespace tl2::tree_stats

