#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
struct MyBoxedArray {
	std::array<int32_t, 2> data{};

	std::string_view tl_name() const { return "myBoxedArray"; }
	uint32_t tl_tag() const { return 0x288f64f0; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

} // namespace tl2

