#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace unique { 
struct Get {
	std::string key;

	std::string_view tl_name() const { return "unique.get"; }
	uint32_t tl_tag() const { return 0xce89bbf2; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::optional<int32_t> & result);
	bool write_result(::basictl::tl_ostream & s, std::optional<int32_t> & result);
};

}} // namespace tl2::unique

