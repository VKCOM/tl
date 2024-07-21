#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "cases.testUnion2.hpp"
#include "cases.testUnion1.hpp"


namespace tl2 { namespace cases { 
struct TestUnion {
	std::variant<::tl2::cases::TestUnion1, ::tl2::cases::TestUnion2> value;

	bool is_1() const { return value.index() == 0; }
	bool is_2() const { return value.index() == 1; }


	std::string_view tl_name() const;
	uint32_t tl_tag() const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases

