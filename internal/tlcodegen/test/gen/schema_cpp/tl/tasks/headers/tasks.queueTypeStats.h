#pragma once

#include "../../../basics/basictl.h"
#include "../types/tasks.queueTypeStats.h"

namespace tl2 { namespace details { 

void TasksQueueTypeStatsReset(::tl2::tasks::QueueTypeStats& item);

bool TasksQueueTypeStatsWriteJSON(std::ostream& s, const ::tl2::tasks::QueueTypeStats& item);
bool TasksQueueTypeStatsRead(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeStats& item);
bool TasksQueueTypeStatsWrite(::basictl::tl_ostream & s, const ::tl2::tasks::QueueTypeStats& item);
bool TasksQueueTypeStatsReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeStats& item);
bool TasksQueueTypeStatsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::QueueTypeStats& item);

}} // namespace tl2::details

