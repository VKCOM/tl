// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "ab/types/ab.counterChangeRequestPeriodsOne.h"

namespace tlgen { namespace details { 

void AbCounterChangeRequestPeriodsOneReset(::tlgen::ab::CounterChangeRequestPeriodsOne& item) noexcept;

bool AbCounterChangeRequestPeriodsOneWriteJSON(std::ostream& s, const ::tlgen::ab::CounterChangeRequestPeriodsOne& item) noexcept;
bool AbCounterChangeRequestPeriodsOneRead(::tlgen::basictl::tl_istream & s, ::tlgen::ab::CounterChangeRequestPeriodsOne& item) noexcept; 
bool AbCounterChangeRequestPeriodsOneWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::ab::CounterChangeRequestPeriodsOne& item) noexcept;
bool AbCounterChangeRequestPeriodsOneReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::ab::CounterChangeRequestPeriodsOne& item);
bool AbCounterChangeRequestPeriodsOneWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::ab::CounterChangeRequestPeriodsOne& item);

}} // namespace tlgen::details

