#pragma once

#include "../../../basics/basictl.h"
#include "cases.TestUnion.h"


namespace tl2 { namespace cases { 
struct TestUnionContainer {
	::tl2::cases::TestUnion value;

	std::string_view tl_name() const { return "cases.testUnionContainer"; }
	uint32_t tl_tag() const { return 0x4497a381; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestUnionContainer& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases

