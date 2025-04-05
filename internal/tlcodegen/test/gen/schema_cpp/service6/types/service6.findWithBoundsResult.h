#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { namespace service6 { 
struct FindWithBoundsResult {
	int32_t x = 0;

	std::string_view tl_name() const { return "service6.findWithBoundsResult"; }
	uint32_t tl_tag() const { return 0x3ded850a; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const FindWithBoundsResult& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service6

