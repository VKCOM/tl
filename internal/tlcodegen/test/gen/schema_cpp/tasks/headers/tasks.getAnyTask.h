#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/tasks.getAnyTask.h"
#include "../types/tasks.taskInfo.h"

namespace tl2 { namespace details { 

void TasksGetAnyTaskReset(::tl2::tasks::GetAnyTask& item);

bool TasksGetAnyTaskWriteJSON(std::ostream& s, const ::tl2::tasks::GetAnyTask& item);
bool TasksGetAnyTaskRead(::basictl::tl_istream & s, ::tl2::tasks::GetAnyTask& item);
bool TasksGetAnyTaskWrite(::basictl::tl_ostream & s, const ::tl2::tasks::GetAnyTask& item);
bool TasksGetAnyTaskReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::GetAnyTask& item);
bool TasksGetAnyTaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::GetAnyTask& item);

bool TasksGetAnyTaskReadResult(::basictl::tl_istream & s, ::tl2::tasks::GetAnyTask& item, std::optional<::tl2::tasks::TaskInfo>& result);
bool TasksGetAnyTaskWriteResult(::basictl::tl_ostream & s, ::tl2::tasks::GetAnyTask& item, std::optional<::tl2::tasks::TaskInfo>& result);
		
}} // namespace tl2::details

