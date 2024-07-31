#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "cases.testOutFieldMask.hpp"


namespace tl2 { namespace cases { 
struct TestOutFieldMaskContainer {
	uint32_t f = 0;
	::tl2::cases::TestOutFieldMask inner{};

	std::string_view tl_name() const { return "cases.testOutFieldMaskContainer"; }
	uint32_t tl_tag() const { return 0x1850ffe4; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases

