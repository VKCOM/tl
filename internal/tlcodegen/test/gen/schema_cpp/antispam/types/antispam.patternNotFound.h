#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { namespace antispam { 
struct PatternNotFound {

	std::string_view tl_name() const { return "antispam.patternNotFound"; }
	uint32_t tl_tag() const { return 0x2c22e225; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const PatternNotFound& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::antispam

