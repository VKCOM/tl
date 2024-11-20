#pragma once

#include "../../../basics/basictl.h"


namespace tl2 { namespace service1 { 
struct StrvalueWithTime {
	std::string value;
	int32_t flags = 0;
	int32_t modificationTime = 0;

	std::string_view tl_name() const { return "service1.strvalueWithTime"; }
	uint32_t tl_tag() const { return 0x98b1a484; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const StrvalueWithTime& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service1

