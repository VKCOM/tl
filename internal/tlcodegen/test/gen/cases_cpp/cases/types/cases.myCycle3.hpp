#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace cases { 
struct MyCycle1;
}} // namespace tl2::cases

namespace tl2 { namespace cases { 
struct MyCycle3 {
	uint32_t fields_mask = 0;
	std::shared_ptr<::tl2::cases::MyCycle1> a{};

	~MyCycle3() {}

	std::string_view tl_name() const { return "cases.myCycle3"; }
	uint32_t tl_tag() const { return 0x7624f86b; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases

