#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/tasks.cronTime.hpp"

namespace tl2 { namespace details { 

void TasksCronTimeReset(::tl2::tasks::CronTime& item);
bool TasksCronTimeRead(::basictl::tl_istream & s, ::tl2::tasks::CronTime& item);
bool TasksCronTimeWrite(::basictl::tl_ostream & s, const ::tl2::tasks::CronTime& item);
bool TasksCronTimeReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::CronTime& item);
bool TasksCronTimeWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::CronTime& item);

}} // namespace tl2::details

