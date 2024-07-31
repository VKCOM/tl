#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/tasks.queueTypeInfo.hpp"

namespace tl2 { namespace details { 

void BuiltinVectorTasksQueueTypeInfoReset(std::vector<::tl2::tasks::QueueTypeInfo>& item);

bool BuiltinVectorTasksQueueTypeInfoWriteJSON(std::ostream & s, const std::vector<::tl2::tasks::QueueTypeInfo>& item);
bool BuiltinVectorTasksQueueTypeInfoRead(::basictl::tl_istream & s, std::vector<::tl2::tasks::QueueTypeInfo>& item);
bool BuiltinVectorTasksQueueTypeInfoWrite(::basictl::tl_ostream & s, const std::vector<::tl2::tasks::QueueTypeInfo>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TasksQueueTypeInfoReset(::tl2::tasks::QueueTypeInfo& item);

bool TasksQueueTypeInfoWriteJSON(std::ostream& s, const ::tl2::tasks::QueueTypeInfo& item);
bool TasksQueueTypeInfoRead(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeInfo& item);
bool TasksQueueTypeInfoWrite(::basictl::tl_ostream & s, const ::tl2::tasks::QueueTypeInfo& item);
bool TasksQueueTypeInfoReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeInfo& item);
bool TasksQueueTypeInfoWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::QueueTypeInfo& item);

}} // namespace tl2::details

