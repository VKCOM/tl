#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service3.limits.hpp"


namespace tl2 { namespace service3 { 
struct GetLimits {

	std::string_view tl_name() const { return "service3.getLimits"; }
	uint32_t tl_tag() const { return 0xeb399467; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::service3::Limits & result);
	bool write_result(::basictl::tl_ostream & s, ::tl2::service3::Limits & result);
};

}} // namespace tl2::service3

