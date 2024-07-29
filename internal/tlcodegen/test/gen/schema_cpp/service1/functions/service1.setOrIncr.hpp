#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service1.Value.hpp"


namespace tl2 { namespace service1 { 
struct SetOrIncr {
	std::string key;
	int32_t flags = 0;
	int32_t delay = 0;
	int64_t value = 0;

	std::string_view tl_name() const { return "service1.setOrIncr"; }
	uint32_t tl_tag() const { return 0x772e390d; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::service1::Value & result);
	bool write_result(::basictl::tl_ostream & s, ::tl2::service1::Value & result);
};

}} // namespace tl2::service1

