#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { namespace cases { 
struct TestUnion2 {
	std::string value;

	std::string_view tl_name() const { return "cases.testUnion2"; }
	uint32_t tl_tag() const { return 0x464f96c4; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestUnion2& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases

