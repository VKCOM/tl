#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
struct FieldConflict3 {
	int32_t x = 0;
	int32_t SetX = 0;

	std::string_view tl_name() const { return "fieldConflict3"; }
	uint32_t tl_tag() const { return 0x2cf6e157; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

} // namespace tl2

