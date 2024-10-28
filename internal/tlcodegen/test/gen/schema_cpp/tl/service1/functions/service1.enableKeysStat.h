#pragma once

#include "../../../basics/basictl.h"


namespace tl2 { namespace service1 { 
struct EnableKeysStat {
	// tl magic for function
	static constexpr uint32_t MAGIC() { return 0x29a7090e; }

	int32_t period = 0;

	std::string_view tl_name() const { return "service1.enableKeysStat"; }
	uint32_t tl_tag() const { return 0x29a7090e; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, bool & result);
	bool write_result(::basictl::tl_ostream & s, bool & result);

	friend std::ostream& operator<<(std::ostream& s, const EnableKeysStat& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service1
