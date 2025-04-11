#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/tasks.cronTask.h"

namespace tl2 { namespace details { 

void TasksCronTaskReset(::tl2::tasks::CronTask& item) noexcept;

bool TasksCronTaskWriteJSON(std::ostream& s, const ::tl2::tasks::CronTask& item) noexcept;
bool TasksCronTaskRead(::basictl::tl_istream & s, ::tl2::tasks::CronTask& item) noexcept; 
bool TasksCronTaskWrite(::basictl::tl_ostream & s, const ::tl2::tasks::CronTask& item) noexcept;
bool TasksCronTaskReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::CronTask& item);
bool TasksCronTaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::CronTask& item);

}} // namespace tl2::details

