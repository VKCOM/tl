#pragma once

#include "../../../basics/basictl.h"


namespace tl2 { namespace benchmarks { 
struct VrutoytopLevelUnionEmpty {

	std::string_view tl_name() const { return "benchmarks.vrutoytopLevelUnionEmpty"; }
	uint32_t tl_tag() const { return 0xce27c770; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const VrutoytopLevelUnionEmpty& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::benchmarks

