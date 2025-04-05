#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { namespace service5 { 
struct EmptyOutput {

	std::string_view tl_name() const { return "service5.emptyOutput"; }
	uint32_t tl_tag() const { return 0x11e46879; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const EmptyOutput& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service5

