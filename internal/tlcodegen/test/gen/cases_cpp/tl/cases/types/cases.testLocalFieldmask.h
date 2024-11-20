#pragma once

#include "../../../basics/basictl.h"
#include "../../__common_namespace/types/true.h"


namespace tl2 { namespace cases { 
struct TestLocalFieldmask {
	uint32_t f1 = 0;
	uint32_t f2 = 0;
	::tl2::True f3{};
	::tl2::True f4{};

	std::string_view tl_name() const { return "cases.testLocalFieldmask"; }
	uint32_t tl_tag() const { return 0xf68fd3f9; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestLocalFieldmask& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases

