#pragma once

#include "../../../basics/basictl.h"


namespace tl2 { namespace service3 { 
struct ProductStatsOld {
	int32_t type = 0;
	int32_t count_new = 0;
	int32_t count_total = 0;
	int32_t count_scheduled = 0;
	int32_t next_scheduled_at = 0;

	std::string_view tl_name() const { return "service3.productStatsOld"; }
	uint32_t tl_tag() const { return 0x6319810b; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const ProductStatsOld& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service3

