#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/tasks.cronTaskWithId.h"

namespace tl2 { namespace details { 

void TasksCronTaskWithIdReset(::tl2::tasks::CronTaskWithId& item);

bool TasksCronTaskWithIdWriteJSON(std::ostream& s, const ::tl2::tasks::CronTaskWithId& item);
bool TasksCronTaskWithIdRead(::basictl::tl_istream & s, ::tl2::tasks::CronTaskWithId& item);
bool TasksCronTaskWithIdWrite(::basictl::tl_ostream & s, const ::tl2::tasks::CronTaskWithId& item);
bool TasksCronTaskWithIdReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::CronTaskWithId& item);
bool TasksCronTaskWithIdWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::CronTaskWithId& item);

}} // namespace tl2::details

