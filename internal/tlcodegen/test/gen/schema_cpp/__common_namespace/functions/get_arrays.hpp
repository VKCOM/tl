#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
struct Get_arrays {
	uint32_t n = 0;
	std::vector<int32_t> a;
	std::array<int32_t, 5> b{};

	std::string_view tl_name() const { return "get_arrays"; }
	uint32_t tl_tag() const { return 0x90658cdb; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::array<int32_t, 5> & result);
	bool write_result(::basictl::tl_ostream & s, std::array<int32_t, 5> & result);
};

} // namespace tl2

