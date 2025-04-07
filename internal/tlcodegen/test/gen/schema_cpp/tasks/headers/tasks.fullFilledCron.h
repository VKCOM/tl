#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/tasks.fullFilledCron.h"

namespace tl2 { namespace details { 

void TasksFullFilledCronReset(::tl2::tasks::FullFilledCron& item);

bool TasksFullFilledCronWriteJSON(std::ostream& s, const ::tl2::tasks::FullFilledCron& item);
bool TasksFullFilledCronRead(::basictl::tl_istream & s, ::tl2::tasks::FullFilledCron& item);
bool TasksFullFilledCronWrite(::basictl::tl_ostream & s, const ::tl2::tasks::FullFilledCron& item);
bool TasksFullFilledCronReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::FullFilledCron& item);
bool TasksFullFilledCronWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::FullFilledCron& item);

}} // namespace tl2::details

