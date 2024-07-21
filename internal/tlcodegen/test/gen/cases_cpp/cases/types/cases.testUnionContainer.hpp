#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "cases.TestUnion.hpp"


namespace tl2 { namespace cases { 
struct TestUnionContainer {
	::tl2::cases::TestUnion value;

	std::string_view tl_name() const { return "cases.testUnionContainer"; }
	uint32_t tl_tag() const { return 0x4497a381; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases

