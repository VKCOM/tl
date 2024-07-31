#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/MyValue.hpp"


namespace tl2 { 
struct GetMyValue {
	::tl2::MyValue x;

	std::string_view tl_name() const { return "getMyValue"; }
	uint32_t tl_tag() const { return 0xb3df27fe; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::MyValue & result);
	bool write_result(::basictl::tl_ostream & s, ::tl2::MyValue & result);
};

} // namespace tl2

