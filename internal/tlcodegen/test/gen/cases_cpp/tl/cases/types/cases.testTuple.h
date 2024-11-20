#pragma once

#include "../../../basics/basictl.h"


namespace tl2 { namespace cases { 
struct TestTuple {
	std::array<int32_t, 4> tpl{};

	std::string_view tl_name() const { return "cases.testTuple"; }
	uint32_t tl_tag() const { return 0x4b9caf8f; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestTuple& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases

