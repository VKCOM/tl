#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/tasks.queueStats.hpp"

namespace tl2 { namespace details { 

void TasksQueueStatsReset(::tl2::tasks::QueueStats& item);

bool TasksQueueStatsWriteJSON(std::ostream& s, const ::tl2::tasks::QueueStats& item, uint32_t nat_fields_mask);
bool TasksQueueStatsRead(::basictl::tl_istream & s, ::tl2::tasks::QueueStats& item, uint32_t nat_fields_mask);
bool TasksQueueStatsWrite(::basictl::tl_ostream & s, const ::tl2::tasks::QueueStats& item, uint32_t nat_fields_mask);
bool TasksQueueStatsReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::QueueStats& item, uint32_t nat_fields_mask);
bool TasksQueueStatsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::QueueStats& item, uint32_t nat_fields_mask);

}} // namespace tl2::details

