#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/tasks.queueStats.h"

namespace tl2 { namespace details { 

void TasksQueueStatsReset(::tl2::tasks::QueueStats& item) noexcept;

bool TasksQueueStatsWriteJSON(std::ostream& s, const ::tl2::tasks::QueueStats& item, uint32_t nat_fields_mask) noexcept;
bool TasksQueueStatsRead(::basictl::tl_istream & s, ::tl2::tasks::QueueStats& item, uint32_t nat_fields_mask) noexcept; 
bool TasksQueueStatsWrite(::basictl::tl_ostream & s, const ::tl2::tasks::QueueStats& item, uint32_t nat_fields_mask) noexcept;
bool TasksQueueStatsReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::QueueStats& item, uint32_t nat_fields_mask);
bool TasksQueueStatsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::QueueStats& item, uint32_t nat_fields_mask);

}} // namespace tl2::details

