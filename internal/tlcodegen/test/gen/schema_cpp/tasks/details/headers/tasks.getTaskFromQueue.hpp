#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../functions/tasks.getTaskFromQueue.hpp"
#include "../../types/tasks.taskInfo.hpp"

namespace tl2 { namespace details { 

void TasksGetTaskFromQueueReset(::tl2::tasks::GetTaskFromQueue& item);
bool TasksGetTaskFromQueueRead(::basictl::tl_istream & s, ::tl2::tasks::GetTaskFromQueue& item);
bool TasksGetTaskFromQueueWrite(::basictl::tl_ostream & s, const ::tl2::tasks::GetTaskFromQueue& item);
bool TasksGetTaskFromQueueReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::GetTaskFromQueue& item);
bool TasksGetTaskFromQueueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::GetTaskFromQueue& item);

bool TasksGetTaskFromQueueReadResult(::basictl::tl_istream & s, ::tl2::tasks::GetTaskFromQueue& item, std::optional<::tl2::tasks::TaskInfo>& result);
bool TasksGetTaskFromQueueWriteResult(::basictl::tl_ostream & s, ::tl2::tasks::GetTaskFromQueue& item, std::optional<::tl2::tasks::TaskInfo>& result);
		
}} // namespace tl2::details

