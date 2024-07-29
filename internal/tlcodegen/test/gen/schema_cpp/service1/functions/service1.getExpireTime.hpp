#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace service1 { 
struct GetExpireTime {
	std::string key;

	std::string_view tl_name() const { return "service1.getExpireTime"; }
	uint32_t tl_tag() const { return 0x5a731070; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::optional<int32_t> & result);
	bool write_result(::basictl::tl_ostream & s, std::optional<int32_t> & result);
};

}} // namespace tl2::service1

