#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { 
struct GetFloat {
	float x = 0;

	std::string_view tl_name() const { return "getFloat"; }
	uint32_t tl_tag() const { return 0x25a7bc68; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, float & result);
	bool write_result(::basictl::tl_ostream & s, float & result);

	friend std::ostream& operator<<(std::ostream& s, const GetFloat& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

