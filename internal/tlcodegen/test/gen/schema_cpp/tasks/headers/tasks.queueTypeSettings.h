#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/tasks.queueTypeSettings.h"

namespace tl2 { namespace details { 

void TasksQueueTypeSettingsReset(::tl2::tasks::QueueTypeSettings& item);

bool TasksQueueTypeSettingsWriteJSON(std::ostream& s, const ::tl2::tasks::QueueTypeSettings& item);
bool TasksQueueTypeSettingsRead(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeSettings& item);
bool TasksQueueTypeSettingsWrite(::basictl::tl_ostream & s, const ::tl2::tasks::QueueTypeSettings& item);
bool TasksQueueTypeSettingsReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeSettings& item);
bool TasksQueueTypeSettingsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::QueueTypeSettings& item);

}} // namespace tl2::details

