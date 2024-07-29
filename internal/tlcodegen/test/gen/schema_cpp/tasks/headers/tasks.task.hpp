#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/tasks.task.hpp"

namespace tl2 { namespace details { 

void TasksTaskReset(::tl2::tasks::Task& item);
bool TasksTaskRead(::basictl::tl_istream & s, ::tl2::tasks::Task& item);
bool TasksTaskWrite(::basictl::tl_ostream & s, const ::tl2::tasks::Task& item);
bool TasksTaskReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::Task& item);
bool TasksTaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::Task& item);

}} // namespace tl2::details

