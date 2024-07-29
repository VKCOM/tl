#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "tasks.TaskStatusItems.hpp"


namespace tl2 { namespace tasks { 
struct TaskStatus {
	std::variant<::tl2::tasks::TaskStatusNotCurrentlyInEngine, ::tl2::tasks::TaskStatusScheduled, ::tl2::tasks::TaskStatusWaiting, ::tl2::tasks::TaskStatusInProgress> value;

	bool is_NotCurrentlyInEngine() const { return value.index() == 0; }
	bool is_Scheduled() const { return value.index() == 1; }
	bool is_Waiting() const { return value.index() == 2; }
	bool is_InProgress() const { return value.index() == 3; }

	void set_NotCurrentlyInEngine() { value.emplace<0>(); }
	void set_Scheduled() { value.emplace<1>(); }
	void set_Waiting() { value.emplace<2>(); }
	void set_InProgress() { value.emplace<3>(); }

	std::string_view tl_name() const;
	uint32_t tl_tag() const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::tasks

