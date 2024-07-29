#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/tasks.taskInfo.hpp"


namespace tl2 { namespace tasks { 
struct GetAnyTask {

	std::string_view tl_name() const { return "tasks.getAnyTask"; }
	uint32_t tl_tag() const { return 0x4a9c7dbb; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::optional<::tl2::tasks::TaskInfo> & result);
	bool write_result(::basictl::tl_ostream & s, std::optional<::tl2::tasks::TaskInfo> & result);
};

}} // namespace tl2::tasks

