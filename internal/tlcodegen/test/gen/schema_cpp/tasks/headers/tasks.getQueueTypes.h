#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "tasks/functions/tasks.getQueueTypes.h"
#include "tasks/types/tasks.queueTypeInfo.h"

namespace tl2 { namespace details { 

void TasksGetQueueTypesReset(::tl2::tasks::GetQueueTypes& item) noexcept;

bool TasksGetQueueTypesWriteJSON(std::ostream& s, const ::tl2::tasks::GetQueueTypes& item) noexcept;
bool TasksGetQueueTypesRead(::basictl::tl_istream & s, ::tl2::tasks::GetQueueTypes& item) noexcept; 
bool TasksGetQueueTypesWrite(::basictl::tl_ostream & s, const ::tl2::tasks::GetQueueTypes& item) noexcept;
bool TasksGetQueueTypesReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::GetQueueTypes& item);
bool TasksGetQueueTypesWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::GetQueueTypes& item);

bool TasksGetQueueTypesReadResult(::basictl::tl_istream & s, ::tl2::tasks::GetQueueTypes& item, std::vector<::tl2::tasks::QueueTypeInfo>& result);
bool TasksGetQueueTypesWriteResult(::basictl::tl_ostream & s, ::tl2::tasks::GetQueueTypes& item, std::vector<::tl2::tasks::QueueTypeInfo>& result);
		
}} // namespace tl2::details

