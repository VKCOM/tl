#pragma once

#include "../../../basics/basictl.h"


namespace tl2 { namespace service1 { 
struct DisableExpiration {
	// tl magic for function
	static constexpr uint32_t MAGIC() { return 0xf1c39c2d; }

	std::string prefix;

	std::string_view tl_name() const { return "service1.disableExpiration"; }
	uint32_t tl_tag() const { return 0xf1c39c2d; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, bool & result);
	bool write_result(::basictl::tl_ostream & s, bool & result);

	friend std::ostream& operator<<(std::ostream& s, const DisableExpiration& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service1

