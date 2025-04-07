#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/tasks.cronTime.h"

namespace tl2 { namespace details { 

void TasksCronTimeReset(::tl2::tasks::CronTime& item);

bool TasksCronTimeWriteJSON(std::ostream& s, const ::tl2::tasks::CronTime& item);
bool TasksCronTimeRead(::basictl::tl_istream & s, ::tl2::tasks::CronTime& item);
bool TasksCronTimeWrite(::basictl::tl_ostream & s, const ::tl2::tasks::CronTime& item);
bool TasksCronTimeReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::CronTime& item);
bool TasksCronTimeWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::CronTime& item);

}} // namespace tl2::details

