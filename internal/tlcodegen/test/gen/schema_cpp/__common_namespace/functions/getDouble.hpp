#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
struct GetDouble {
	double x = 0;

	std::string_view tl_name() const { return "getDouble"; }
	uint32_t tl_tag() const { return 0x39711d7b; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, double & result);
	bool write_result(::basictl::tl_ostream & s, double & result);
};

} // namespace tl2

