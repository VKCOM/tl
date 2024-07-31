#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
struct BoolStat {
	int32_t statTrue = 0;
	int32_t statFalse = 0;
	int32_t statUnknown = 0;

	std::string_view tl_name() const { return "boolStat"; }
	uint32_t tl_tag() const { return 0x92cbcbfa; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const BoolStat& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

