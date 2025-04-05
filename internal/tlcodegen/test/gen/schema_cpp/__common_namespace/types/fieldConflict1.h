#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { 
struct FieldConflict1 {
	int32_t x = 0;
	int32_t set_x = 0;

	std::string_view tl_name() const { return "fieldConflict1"; }
	uint32_t tl_tag() const { return 0xf314bd09; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const FieldConflict1& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

