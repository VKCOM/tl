#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/myDictOfInt.hpp"


namespace tl2 { 
struct GetMyDictOfInt {
	::tl2::MyDictOfInt x{};

	std::string_view tl_name() const { return "getMyDictOfInt"; }
	uint32_t tl_tag() const { return 0x166f962c; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::MyDictOfInt & result);
	bool write_result(::basictl::tl_ostream & s, ::tl2::MyDictOfInt & result);
};

} // namespace tl2

