#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
struct FieldConflict2 {
	int32_t x = 0;
	int32_t setX = 0;

	std::string_view tl_name() const { return "fieldConflict2"; }
	uint32_t tl_tag() const { return 0x1bba76b8; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const FieldConflict2& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

