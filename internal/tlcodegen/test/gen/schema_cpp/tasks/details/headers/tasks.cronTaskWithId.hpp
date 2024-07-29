#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../types/tasks.cronTaskWithId.hpp"

namespace tl2 { namespace details { 

void TasksCronTaskWithIdReset(::tl2::tasks::CronTaskWithId& item);
bool TasksCronTaskWithIdRead(::basictl::tl_istream & s, ::tl2::tasks::CronTaskWithId& item);
bool TasksCronTaskWithIdWrite(::basictl::tl_ostream & s, const ::tl2::tasks::CronTaskWithId& item);
bool TasksCronTaskWithIdReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::CronTaskWithId& item);
bool TasksCronTaskWithIdWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::CronTaskWithId& item);

}} // namespace tl2::details

