#pragma once

#include "../../../basics/basictl.h"


namespace tl2 { namespace service1 { 
struct GetKeysStatPeriods {
	// tl magic for function
	static constexpr uint32_t MAGIC() { return 0x8cdf39e3; }


	std::string_view tl_name() const { return "service1.getKeysStatPeriods"; }
	uint32_t tl_tag() const { return 0x8cdf39e3; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::vector<int32_t> & result);
	bool write_result(::basictl::tl_ostream & s, std::vector<int32_t> & result);

	friend std::ostream& operator<<(std::ostream& s, const GetKeysStatPeriods& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service1

