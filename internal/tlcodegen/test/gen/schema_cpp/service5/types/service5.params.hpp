#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace service5 { 
struct Params {
	uint32_t fields_mask = 0;
	int32_t max_execution_speed = 0;
	int32_t max_execution_speed_bytes = 0;

	std::string_view tl_name() const { return "service5.params"; }
	uint32_t tl_tag() const { return 0x12ae5cb5; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::service5

