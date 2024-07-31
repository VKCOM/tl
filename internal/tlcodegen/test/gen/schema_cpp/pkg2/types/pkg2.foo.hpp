#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace pkg2 { 
struct Foo {
	int32_t x = 0;

	std::string_view tl_name() const { return "pkg2.foo"; }
	uint32_t tl_tag() const { return 0xe144703d; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::pkg2

