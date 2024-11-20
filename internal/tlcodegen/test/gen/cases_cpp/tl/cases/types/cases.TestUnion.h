#pragma once

#include "../../../basics/basictl.h"
#include "cases.testUnion2.h"
#include "cases.testUnion1.h"


namespace tl2 { namespace cases { 
struct TestUnion {
	std::variant<::tl2::cases::TestUnion1, ::tl2::cases::TestUnion2> value;

	bool is_1() const { return value.index() == 0; }
	bool is_2() const { return value.index() == 1; }


	std::string_view tl_name() const;
	uint32_t tl_tag() const;

	bool write_json(std::ostream& s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases

