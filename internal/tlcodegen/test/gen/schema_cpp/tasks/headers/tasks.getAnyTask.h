#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "tasks/functions/tasks.getAnyTask.h"
#include "tasks/types/tasks.taskInfo.h"

namespace tl2 { namespace details { 

void TasksGetAnyTaskReset(::tl2::tasks::GetAnyTask& item) noexcept;

bool TasksGetAnyTaskWriteJSON(std::ostream& s, const ::tl2::tasks::GetAnyTask& item) noexcept;
bool TasksGetAnyTaskRead(::basictl::tl_istream & s, ::tl2::tasks::GetAnyTask& item) noexcept; 
bool TasksGetAnyTaskWrite(::basictl::tl_ostream & s, const ::tl2::tasks::GetAnyTask& item) noexcept;
bool TasksGetAnyTaskReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::GetAnyTask& item);
bool TasksGetAnyTaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::GetAnyTask& item);

bool TasksGetAnyTaskReadResult(::basictl::tl_istream & s, ::tl2::tasks::GetAnyTask& item, std::optional<::tl2::tasks::TaskInfo>& result);
bool TasksGetAnyTaskWriteResult(::basictl::tl_ostream & s, ::tl2::tasks::GetAnyTask& item, std::optional<::tl2::tasks::TaskInfo>& result);
		
}} // namespace tl2::details

