#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "tasks.cronTask.hpp"


namespace tl2 { namespace tasks { 
struct CronTaskWithId {
	int32_t id = 0;
	int32_t next_time = 0;
	::tl2::tasks::CronTask task{};

	std::string_view tl_name() const { return "tasks.cronTaskWithId"; }
	uint32_t tl_tag() const { return 0x3a958001; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const CronTaskWithId& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::tasks

