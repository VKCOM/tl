#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace service1 { 
struct LongvalueWithTime {
	int64_t value = 0;
	int32_t flags = 0;
	int32_t modificationTime = 0;

	std::string_view tl_name() const { return "service1.longvalueWithTime"; }
	uint32_t tl_tag() const { return 0xa04606ec; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::service1

