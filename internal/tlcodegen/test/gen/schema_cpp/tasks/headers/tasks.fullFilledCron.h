// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "tasks/types/tasks.fullFilledCron.h"

namespace tlgen { namespace details { 

void TasksFullFilledCronReset(::tlgen::tasks::FullFilledCron& item) noexcept;

bool TasksFullFilledCronWriteJSON(std::ostream& s, const ::tlgen::tasks::FullFilledCron& item) noexcept;
bool TasksFullFilledCronRead(::tlgen::basictl::tl_istream & s, ::tlgen::tasks::FullFilledCron& item) noexcept; 
bool TasksFullFilledCronWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::tasks::FullFilledCron& item) noexcept;
bool TasksFullFilledCronReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::tasks::FullFilledCron& item);
bool TasksFullFilledCronWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::tasks::FullFilledCron& item);

}} // namespace tlgen::details

