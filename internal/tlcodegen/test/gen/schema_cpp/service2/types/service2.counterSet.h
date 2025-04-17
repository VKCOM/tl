#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"


namespace tl2 { namespace service2 { 
struct CounterSet {
	std::vector<int32_t> intCounters;
	std::vector<double> floatCounters;

	std::string_view tl_name() const { return "service2.counterSet"; }
	uint32_t tl_tag() const { return 0xf5403fd9; }

	bool write_json(std::ostream& s, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum)const;

	bool read(::basictl::tl_istream & s, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum) noexcept;
	bool write(::basictl::tl_ostream & s, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum);
	void write_or_throw(::basictl::tl_throwable_ostream & s, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum)const;

	bool read_boxed(::basictl::tl_istream & s, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum) noexcept;
	bool write_boxed(::basictl::tl_ostream & s, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum)const;
};

}} // namespace tl2::service2

