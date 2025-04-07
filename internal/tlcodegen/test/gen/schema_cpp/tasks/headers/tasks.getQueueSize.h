#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/tasks.getQueueSize.h"
#include "../types/tasks.queueStats.h"

namespace tl2 { namespace details { 

void TasksGetQueueSizeReset(::tl2::tasks::GetQueueSize& item);

bool TasksGetQueueSizeWriteJSON(std::ostream& s, const ::tl2::tasks::GetQueueSize& item);
bool TasksGetQueueSizeRead(::basictl::tl_istream & s, ::tl2::tasks::GetQueueSize& item);
bool TasksGetQueueSizeWrite(::basictl::tl_ostream & s, const ::tl2::tasks::GetQueueSize& item);
bool TasksGetQueueSizeReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::GetQueueSize& item);
bool TasksGetQueueSizeWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::GetQueueSize& item);

bool TasksGetQueueSizeReadResult(::basictl::tl_istream & s, ::tl2::tasks::GetQueueSize& item, ::tl2::tasks::QueueStats& result);
bool TasksGetQueueSizeWriteResult(::basictl::tl_ostream & s, ::tl2::tasks::GetQueueSize& item, ::tl2::tasks::QueueStats& result);
		
}} // namespace tl2::details

