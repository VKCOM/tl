#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/tasks.TaskStatus.h"

namespace tl2 { namespace details { 

void TasksTaskStatusReset(::tl2::tasks::TaskStatus& item);

bool TasksTaskStatusWriteJSON(std::ostream & s, const ::tl2::tasks::TaskStatus& item);
bool TasksTaskStatusReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::TaskStatus& item);
bool TasksTaskStatusWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatus& item);

}} // namespace tl2::details

