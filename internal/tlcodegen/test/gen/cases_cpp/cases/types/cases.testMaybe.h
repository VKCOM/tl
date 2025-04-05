#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { namespace cases { 
struct TestMaybe {
	std::optional<int32_t> value;

	std::string_view tl_name() const { return "cases.testMaybe"; }
	uint32_t tl_tag() const { return 0xd6602613; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestMaybe& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases

