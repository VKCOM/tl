#pragma once

#include "../../../basics/basictl.h"


namespace tl2 { namespace cases_bytes { 
struct TestArray {
	uint32_t n = 0;
	std::vector<std::string> arr;

	std::string_view tl_name() const { return "cases_bytes.testArray"; }
	uint32_t tl_tag() const { return 0x3762fb81; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestArray& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases_bytes

