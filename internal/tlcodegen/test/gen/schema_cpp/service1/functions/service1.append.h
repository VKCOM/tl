#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { namespace service1 { 
struct Append {
	std::string key;
	int32_t flags = 0;
	int32_t delay = 0;
	std::string suffix;

	std::string_view tl_name() const { return "service1.append"; }
	uint32_t tl_tag() const { return 0x04dec671; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, bool & result);
	bool write_result(::basictl::tl_ostream & s, bool & result);

	friend std::ostream& operator<<(std::ostream& s, const Append& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service1

