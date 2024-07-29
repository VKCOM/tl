#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "integer.hpp"


namespace tl2 { 
struct BenchObject {
	std::vector<int32_t> xs;
	std::vector<::tl2::Integer> ys;

	std::string_view tl_name() const { return "benchObject"; }
	uint32_t tl_tag() const { return 0xb697e865; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

} // namespace tl2

