#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace tasks { 
struct CronTime {
	uint32_t fields_mask = 0;
	std::vector<int32_t> seconds;
	std::vector<int32_t> minutes;
	std::vector<int32_t> hours;
	std::vector<int32_t> days_of_week;
	std::vector<int32_t> days;
	std::vector<int32_t> months;

	std::string_view tl_name() const { return "tasks.cronTime"; }
	uint32_t tl_tag() const { return 0xd4177d7f; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::tasks

