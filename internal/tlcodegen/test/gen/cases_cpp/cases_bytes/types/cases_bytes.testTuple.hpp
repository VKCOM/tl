#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace cases_bytes { 
struct TestTuple {
	std::array<std::string, 4> tpl{};

	std::string_view tl_name() const { return "cases_bytes.testTuple"; }
	uint32_t tl_tag() const { return 0x2dd3bacf; }

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

}} // namespace tl2::cases_bytes

