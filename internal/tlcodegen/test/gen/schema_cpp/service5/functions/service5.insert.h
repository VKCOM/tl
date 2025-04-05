#pragma once

#include "../../basictl/io_streams.h"
#include "../types/service5.Output.h"


namespace tl2 { namespace service5 { 
struct Insert {
	std::string table;
	std::string data;

	std::string_view tl_name() const { return "service5.insert"; }
	uint32_t tl_tag() const { return 0xc911ee2c; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::service5::Output & result);
	bool write_result(::basictl::tl_ostream & s, ::tl2::service5::Output & result);

	friend std::ostream& operator<<(std::ostream& s, const Insert& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service5

