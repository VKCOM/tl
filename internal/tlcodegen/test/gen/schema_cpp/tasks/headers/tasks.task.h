#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "tasks/types/tasks.task.h"

namespace tl2 { namespace details { 

void TasksTaskReset(::tl2::tasks::Task& item) noexcept;

bool TasksTaskWriteJSON(std::ostream& s, const ::tl2::tasks::Task& item) noexcept;
bool TasksTaskRead(::basictl::tl_istream & s, ::tl2::tasks::Task& item) noexcept; 
bool TasksTaskWrite(::basictl::tl_ostream & s, const ::tl2::tasks::Task& item) noexcept;
bool TasksTaskReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::Task& item);
bool TasksTaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::Task& item);

}} // namespace tl2::details

