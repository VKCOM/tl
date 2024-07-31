#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "benchmarks.VrutoyTopLevelUnion.hpp"


namespace tl2 { namespace benchmarks { 
struct VrutoyTopLevelContainer {
	::tl2::benchmarks::VrutoyTopLevelUnion value;

	std::string_view tl_name() const { return "benchmarks.vrutoyTopLevelContainer"; }
	uint32_t tl_tag() const { return 0xfb442ca5; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::benchmarks

