#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/tasks.addTask.h"

namespace tl2 { namespace details { 

void TasksAddTaskReset(::tl2::tasks::AddTask& item) noexcept;

bool TasksAddTaskWriteJSON(std::ostream& s, const ::tl2::tasks::AddTask& item) noexcept;
bool TasksAddTaskRead(::basictl::tl_istream & s, ::tl2::tasks::AddTask& item) noexcept; 
bool TasksAddTaskWrite(::basictl::tl_ostream & s, const ::tl2::tasks::AddTask& item) noexcept;
bool TasksAddTaskReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::AddTask& item);
bool TasksAddTaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::AddTask& item);

bool TasksAddTaskReadResult(::basictl::tl_istream & s, ::tl2::tasks::AddTask& item, int64_t& result);
bool TasksAddTaskWriteResult(::basictl::tl_ostream & s, ::tl2::tasks::AddTask& item, int64_t& result);
		
}} // namespace tl2::details

