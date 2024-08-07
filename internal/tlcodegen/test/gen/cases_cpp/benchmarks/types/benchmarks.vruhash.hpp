#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace benchmarks { 
struct Vruhash {
	int64_t low = 0;
	int64_t high = 0;

	std::string_view tl_name() const { return "benchmarks.vruhash"; }
	uint32_t tl_tag() const { return 0xd31bd0fd; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const Vruhash& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::benchmarks

