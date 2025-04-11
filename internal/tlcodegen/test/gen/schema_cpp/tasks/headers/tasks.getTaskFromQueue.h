#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/tasks.getTaskFromQueue.h"
#include "../types/tasks.taskInfo.h"

namespace tl2 { namespace details { 

void TasksGetTaskFromQueueReset(::tl2::tasks::GetTaskFromQueue& item) noexcept;

bool TasksGetTaskFromQueueWriteJSON(std::ostream& s, const ::tl2::tasks::GetTaskFromQueue& item) noexcept;
bool TasksGetTaskFromQueueRead(::basictl::tl_istream & s, ::tl2::tasks::GetTaskFromQueue& item) noexcept; 
bool TasksGetTaskFromQueueWrite(::basictl::tl_ostream & s, const ::tl2::tasks::GetTaskFromQueue& item) noexcept;
bool TasksGetTaskFromQueueReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::GetTaskFromQueue& item);
bool TasksGetTaskFromQueueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::GetTaskFromQueue& item);

bool TasksGetTaskFromQueueReadResult(::basictl::tl_istream & s, ::tl2::tasks::GetTaskFromQueue& item, std::optional<::tl2::tasks::TaskInfo>& result);
bool TasksGetTaskFromQueueWriteResult(::basictl::tl_ostream & s, ::tl2::tasks::GetTaskFromQueue& item, std::optional<::tl2::tasks::TaskInfo>& result);
		
}} // namespace tl2::details

