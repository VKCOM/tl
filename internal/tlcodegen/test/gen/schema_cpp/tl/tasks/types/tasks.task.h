#pragma once

#include "../../../basics/basictl.h"


namespace tl2 { namespace tasks { 
struct Task {
	uint32_t fields_mask = 0;
	int32_t flags = 0;
	std::vector<int32_t> tag;
	std::string data;
	int64_t id = 0;
	int32_t retries = 0;
	int32_t scheduled_time = 0;
	int32_t deadline = 0;

	std::string_view tl_name() const { return "tasks.task"; }
	uint32_t tl_tag() const { return 0x7c23bc2c; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const Task& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::tasks

