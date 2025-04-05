#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { namespace service6 { 
struct FindResultRow {
	int32_t x = 0;

	std::string_view tl_name() const { return "service6.findResultRow"; }
	uint32_t tl_tag() const { return 0xbd3946e3; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const FindResultRow& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service6

