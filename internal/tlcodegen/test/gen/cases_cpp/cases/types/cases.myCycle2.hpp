#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "cases.myCycle3.hpp"


namespace tl2 { namespace cases { 
struct MyCycle3;
}} // namespace tl2::cases

namespace tl2 { namespace cases { 
struct MyCycle2 {
	uint32_t fields_mask = 0;
	::tl2::cases::MyCycle3 a{};

	std::string_view tl_name() const { return "cases.myCycle2"; }
	uint32_t tl_tag() const { return 0x5444c9a2; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases

