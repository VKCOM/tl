#pragma once

#include "../../../basics/basictl.h"


namespace tl2 { namespace cases { 
struct TestVector {
	std::vector<int32_t> arr;

	std::string_view tl_name() const { return "cases.testVector"; }
	uint32_t tl_tag() const { return 0x4975695c; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestVector& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases

