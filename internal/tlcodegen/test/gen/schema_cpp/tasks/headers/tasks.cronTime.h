// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "tasks/types/tasks.cronTime.h"

namespace tlgen { namespace details { 

void TasksCronTimeReset(::tlgen::tasks::CronTime& item) noexcept;

bool TasksCronTimeWriteJSON(std::ostream& s, const ::tlgen::tasks::CronTime& item) noexcept;
bool TasksCronTimeRead(::tlgen::basictl::tl_istream & s, ::tlgen::tasks::CronTime& item) noexcept; 
bool TasksCronTimeWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::tasks::CronTime& item) noexcept;
bool TasksCronTimeReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::tasks::CronTime& item);
bool TasksCronTimeWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::tasks::CronTime& item);

}} // namespace tlgen::details

