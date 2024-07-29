#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/antispam.PatternFull.hpp"


namespace tl2 { namespace antispam { 
struct GetPattern {
	int32_t id = 0;

	std::string_view tl_name() const { return "antispam.getPattern"; }
	uint32_t tl_tag() const { return 0x3de14136; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::antispam::PatternFull & result);
	bool write_result(::basictl::tl_ostream & s, ::tl2::antispam::PatternFull & result);
};

}} // namespace tl2::antispam

