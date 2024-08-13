#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service2.deltaSet.hpp"
#include "../../__common_namespace/types/true.hpp"


namespace tl2 { namespace service2 { 
struct Set {
	// tl magic for function
	static const uint32_t MAGIC = 0x0d31f63d;

	uint32_t objectIdLength = 0;
	uint32_t intCountersNum = 0;
	uint32_t floatCountersNum = 0;
	std::vector<int32_t> intCounters;
	std::vector<int32_t> floatCounters;
	::tl2::service2::DeltaSet newValues{};

	std::string_view tl_name() const { return "service2.set"; }
	uint32_t tl_tag() const { return 0x0d31f63d; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::True & result);
	bool write_result(::basictl::tl_ostream & s, ::tl2::True & result);

	friend std::ostream& operator<<(std::ostream& s, const Set& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service2

