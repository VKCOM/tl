#pragma once

#include "../../basictl/io_streams.h"
#include "../types/service2.deltaSet.h"
#include "../../__common_namespace/types/true.h"


namespace tl2 { namespace service2 { 
struct Set {
	uint32_t objectIdLength = 0;
	uint32_t intCountersNum = 0;
	uint32_t floatCountersNum = 0;
	std::vector<int32_t> intCounters;
	std::vector<int32_t> floatCounters;
	::tl2::service2::DeltaSet newValues{};

	std::string_view tl_name() const { return "service2.set"; }
	uint32_t tl_tag() const { return 0x0d31f63d; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::True & result) noexcept;
	bool write_result(::basictl::tl_ostream & s, ::tl2::True & result) noexcept;

	void read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::True & result);
	void write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::True & result);

	friend std::ostream& operator<<(std::ostream& s, const Set& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service2

