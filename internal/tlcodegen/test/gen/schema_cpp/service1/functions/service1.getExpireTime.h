#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { namespace service1 { 
struct GetExpireTime {
	std::string key;

	std::string_view tl_name() const { return "service1.getExpireTime"; }
	uint32_t tl_tag() const { return 0x5a731070; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::optional<int32_t> & result);
	bool write_result(::basictl::tl_ostream & s, std::optional<int32_t> & result);

	friend std::ostream& operator<<(std::ostream& s, const GetExpireTime& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service1

