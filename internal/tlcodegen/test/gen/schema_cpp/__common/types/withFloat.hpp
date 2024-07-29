#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
struct WithFloat {
	float x = 0;
	float y = 0;
	float z = 0;

	std::string_view tl_name() const { return "withFloat"; }
	uint32_t tl_tag() const { return 0x071b8685; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

} // namespace tl2

