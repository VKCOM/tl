#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "tasks.task.hpp"


namespace tl2 { namespace tasks { 
struct TaskInfo {
	std::string type_name;
	std::vector<int32_t> queue_id;
	::tl2::tasks::Task task{};

	std::string_view tl_name() const { return "tasks.taskInfo"; }
	uint32_t tl_tag() const { return 0x06f0c6a6; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::tasks

