#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
struct BoxedVector64 {
	std::vector<int64_t> x;

	std::string_view tl_name() const { return "boxedVector64"; }
	uint32_t tl_tag() const { return 0x83659ba8; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::vector<int64_t> & result);
	bool write_result(::basictl::tl_ostream & s, std::vector<int64_t> & result);
};

} // namespace tl2

