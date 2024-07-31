#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../cases/types/cases.TestEnum.hpp"


namespace tl2 { namespace cases_bytes { 
struct TestEnumContainer {
	::tl2::cases::TestEnum value;

	std::string_view tl_name() const { return "cases_bytes.testEnumContainer"; }
	uint32_t tl_tag() const { return 0x32b92037; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases_bytes

