#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "benchmarks.vruposition.hpp"


namespace tl2 { namespace benchmarks { 
struct VrutoytopLevelUnionBig {
	std::vector<::tl2::benchmarks::Vruposition> next_positions;

	std::string_view tl_name() const { return "benchmarks.vrutoytopLevelUnionBig"; }
	uint32_t tl_tag() const { return 0xef556bee; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::benchmarks

