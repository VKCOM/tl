#pragma once

#include "../../basictl/io_streams.h"
#include "tasks.task.h"


namespace tl2 { namespace tasks { 
struct TaskInfo {
	std::string type_name;
	std::vector<int32_t> queue_id;
	::tl2::tasks::Task task{};

	std::string_view tl_name() const { return "tasks.taskInfo"; }
	uint32_t tl_tag() const { return 0x06f0c6a6; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TaskInfo& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::tasks

