#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "benchmarks.vrutoytopLevelUnionEmpty.hpp"
#include "benchmarks.vrutoytopLevelUnionBig.hpp"


namespace tl2 { namespace benchmarks { 
struct VrutoyTopLevelUnion {
	std::variant<::tl2::benchmarks::VrutoytopLevelUnionBig, ::tl2::benchmarks::VrutoytopLevelUnionEmpty> value;

	bool is_Big() const { return value.index() == 0; }
	bool is_Empty() const { return value.index() == 1; }

	void set_Empty() { value.emplace<1>(); }

	std::string_view tl_name() const;
	uint32_t tl_tag() const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::benchmarks

