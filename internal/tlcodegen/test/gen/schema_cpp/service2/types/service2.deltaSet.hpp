#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "service2.objectId.hpp"
#include "service2.counterSet.hpp"


namespace tl2 { namespace service2 { 
struct DeltaSet {
	::tl2::service2::ObjectId id{};
	::tl2::service2::CounterSet counters{};

	std::string_view tl_name() const { return "service2.deltaSet"; }
	uint32_t tl_tag() const { return 0xbf49abc2; }

	bool read(::basictl::tl_istream & s, uint32_t nat_objectIdLength, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum);
	bool write(::basictl::tl_ostream & s, uint32_t nat_objectIdLength, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum)const;

	bool read_boxed(::basictl::tl_istream & s, uint32_t nat_objectIdLength, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum);
	bool write_boxed(::basictl::tl_ostream & s, uint32_t nat_objectIdLength, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum)const;
};

}} // namespace tl2::service2

