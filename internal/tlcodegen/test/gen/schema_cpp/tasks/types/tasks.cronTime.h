#pragma once

#include "../../basictl/io_streams.h"


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

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const CronTime& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::tasks

