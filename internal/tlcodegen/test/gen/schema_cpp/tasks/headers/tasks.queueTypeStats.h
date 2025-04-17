#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "tasks/types/tasks.queueTypeStats.h"

namespace tl2 { namespace details { 

void TasksQueueTypeStatsReset(::tl2::tasks::QueueTypeStats& item) noexcept;

bool TasksQueueTypeStatsWriteJSON(std::ostream& s, const ::tl2::tasks::QueueTypeStats& item) noexcept;
bool TasksQueueTypeStatsRead(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeStats& item) noexcept; 
bool TasksQueueTypeStatsWrite(::basictl::tl_ostream & s, const ::tl2::tasks::QueueTypeStats& item) noexcept;
bool TasksQueueTypeStatsReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeStats& item);
bool TasksQueueTypeStatsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::QueueTypeStats& item);

}} // namespace tl2::details

