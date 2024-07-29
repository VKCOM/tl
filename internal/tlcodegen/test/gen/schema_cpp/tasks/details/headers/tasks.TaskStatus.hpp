#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../types/tasks.TaskStatus.hpp"

namespace tl2 { namespace details { 

void TasksTaskStatusReset(::tl2::tasks::TaskStatus& item);
bool TasksTaskStatusReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::TaskStatus& item);
bool TasksTaskStatusWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatus& item);

}} // namespace tl2::details

