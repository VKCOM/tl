#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace cases { 
struct Replace7plusplus {
	uint32_t N = 0;
	uint32_t M = 0;
	std::vector<std::vector<int32_t>> A;

	std::string_view tl_name() const { return "cases.replace7plusplus"; }
	uint32_t tl_tag() const { return 0xabc39b68; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases

