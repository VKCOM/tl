#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../types/tasks.queueTypeSettings.hpp"

namespace tl2 { namespace details { 

void TasksQueueTypeSettingsReset(::tl2::tasks::QueueTypeSettings& item);
bool TasksQueueTypeSettingsRead(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeSettings& item);
bool TasksQueueTypeSettingsWrite(::basictl::tl_ostream & s, const ::tl2::tasks::QueueTypeSettings& item);
bool TasksQueueTypeSettingsReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeSettings& item);
bool TasksQueueTypeSettingsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::QueueTypeSettings& item);

}} // namespace tl2::details

