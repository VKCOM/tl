#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service1.Value.hpp"
#include "../../__common/types/dictionary.hpp"


namespace tl2 { namespace service1 { 
struct GetWildcardWithFlags {
	std::string prefix;

	std::string_view tl_name() const { return "service1.getWildcardWithFlags"; }
	uint32_t tl_tag() const { return 0x5f6a1f78; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::Dictionary<::tl2::service1::Value> & result);
	bool write_result(::basictl::tl_ostream & s, ::tl2::Dictionary<::tl2::service1::Value> & result);
};

}} // namespace tl2::service1

