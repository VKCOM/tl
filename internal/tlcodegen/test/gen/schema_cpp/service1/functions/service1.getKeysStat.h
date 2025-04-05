#pragma once

#include "../../basictl/io_streams.h"
#include "../types/service1.keysStat.h"


namespace tl2 { namespace service1 { 
struct GetKeysStat {
	int32_t period = 0;

	std::string_view tl_name() const { return "service1.getKeysStat"; }
	uint32_t tl_tag() const { return 0x06cecd58; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::optional<::tl2::service1::KeysStat> & result);
	bool write_result(::basictl::tl_ostream & s, std::optional<::tl2::service1::KeysStat> & result);

	friend std::ostream& operator<<(std::ostream& s, const GetKeysStat& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service1

