#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace service2 { 
struct CounterSet {
	std::vector<int32_t> intCounters;
	std::vector<double> floatCounters;

	std::string_view tl_name() const { return "service2.counterSet"; }
	uint32_t tl_tag() const { return 0xf5403fd9; }

	bool read(::basictl::tl_istream & s, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum);
	bool write(::basictl::tl_ostream & s, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum)const;

	bool read_boxed(::basictl::tl_istream & s, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum);
	bool write_boxed(::basictl::tl_ostream & s, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum)const;
};

}} // namespace tl2::service2

