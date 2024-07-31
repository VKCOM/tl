#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "tree_stats.objectLimitValueLong.hpp"
#include "tree_stats.objectLimitValueDouble.hpp"


namespace tl2 { namespace tree_stats { 
struct ObjectLimitValue {
	std::variant<::tl2::tree_stats::ObjectLimitValueLong, ::tl2::tree_stats::ObjectLimitValueDouble> value;

	bool is_Long() const { return value.index() == 0; }
	bool is_Double() const { return value.index() == 1; }

	void set_Long() { value.emplace<0>(); }

	std::string_view tl_name() const;
	uint32_t tl_tag() const;

	bool write_json(std::ostream& s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::tree_stats

