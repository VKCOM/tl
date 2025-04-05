#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { 
struct WithFloat {
	float x = 0;
	float y = 0;
	float z = 0;

	std::string_view tl_name() const { return "withFloat"; }
	uint32_t tl_tag() const { return 0x071b8685; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const WithFloat& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

