#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace antispam { 
struct PatternFound {
	int32_t ip = 0;
	int32_t uahash = 0;
	int32_t flags = 0;
	int32_t type = 0;
	std::string text;

	std::string_view tl_name() const { return "antispam.patternFound"; }
	uint32_t tl_tag() const { return 0xa7688492; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::antispam

