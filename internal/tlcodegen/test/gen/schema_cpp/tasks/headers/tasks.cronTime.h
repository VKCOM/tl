#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "tasks/types/tasks.cronTime.h"

namespace tl2 { namespace details { 

void TasksCronTimeReset(::tl2::tasks::CronTime& item) noexcept;

bool TasksCronTimeWriteJSON(std::ostream& s, const ::tl2::tasks::CronTime& item) noexcept;
bool TasksCronTimeRead(::basictl::tl_istream & s, ::tl2::tasks::CronTime& item) noexcept; 
bool TasksCronTimeWrite(::basictl::tl_ostream & s, const ::tl2::tasks::CronTime& item) noexcept;
bool TasksCronTimeReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::CronTime& item);
bool TasksCronTimeWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::CronTime& item);

}} // namespace tl2::details

