#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { namespace cases { 
struct Replace7 {
	uint32_t n = 0;
	uint32_t m = 0;
	std::vector<std::vector<int32_t>> a;

	std::string_view tl_name() const { return "cases.replace7"; }
	uint32_t tl_tag() const { return 0x6ccce4be; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const Replace7& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases

