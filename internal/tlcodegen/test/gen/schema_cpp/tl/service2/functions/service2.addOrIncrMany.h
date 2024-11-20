#pragma once

#include "../../../basics/basictl.h"
#include "../types/service2.deltaSet.h"
#include "../types/service2.counterSet.h"


namespace tl2 { namespace service2 { 
struct AddOrIncrMany {
	// tl magic for function
	static constexpr uint32_t MAGIC() { return 0x5aa52489; }

	uint32_t objectIdLength = 0;
	uint32_t intCountersNum = 0;
	uint32_t floatCountersNum = 0;
	uint32_t objectsNum = 0;
	std::vector<int32_t> intCounters;
	std::vector<int32_t> floatCounters;
	std::vector<::tl2::service2::DeltaSet> deltas;

	std::string_view tl_name() const { return "service2.addOrIncrMany"; }
	uint32_t tl_tag() const { return 0x5aa52489; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::vector<::tl2::service2::CounterSet> & result);
	bool write_result(::basictl::tl_ostream & s, std::vector<::tl2::service2::CounterSet> & result);

	friend std::ostream& operator<<(std::ostream& s, const AddOrIncrMany& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service2

