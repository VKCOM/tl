#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "benchmarks.vrutoyPositions.hpp"


namespace tl2 { namespace benchmarks { 
struct VrutoyTopLevelContainerWithDependency {
	uint32_t n = 0;
	::tl2::benchmarks::VrutoyPositions value{};

	std::string_view tl_name() const { return "benchmarks.vrutoyTopLevelContainerWithDependency"; }
	uint32_t tl_tag() const { return 0xc176008e; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const VrutoyTopLevelContainerWithDependency& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::benchmarks

