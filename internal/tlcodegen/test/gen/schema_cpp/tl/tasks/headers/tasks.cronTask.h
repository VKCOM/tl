#pragma once

#include "../../../basics/basictl.h"
#include "../types/tasks.cronTask.h"

namespace tl2 { namespace details { 

void TasksCronTaskReset(::tl2::tasks::CronTask& item);

bool TasksCronTaskWriteJSON(std::ostream& s, const ::tl2::tasks::CronTask& item);
bool TasksCronTaskRead(::basictl::tl_istream & s, ::tl2::tasks::CronTask& item);
bool TasksCronTaskWrite(::basictl::tl_ostream & s, const ::tl2::tasks::CronTask& item);
bool TasksCronTaskReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::CronTask& item);
bool TasksCronTaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::CronTask& item);

}} // namespace tl2::details

