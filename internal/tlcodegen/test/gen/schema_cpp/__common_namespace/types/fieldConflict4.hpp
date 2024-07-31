#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
struct FieldConflict4 {
	int32_t X = 0;
	int32_t SetX = 0;

	std::string_view tl_name() const { return "fieldConflict4"; }
	uint32_t tl_tag() const { return 0xd93c186a; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

} // namespace tl2

