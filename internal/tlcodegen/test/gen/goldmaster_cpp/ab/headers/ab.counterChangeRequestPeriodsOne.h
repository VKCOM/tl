// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "ab/types/ab.counterChangeRequestPeriodsOne.h"

namespace tl2 { namespace details { 

void AbCounterChangeRequestPeriodsOneReset(::tl2::ab::CounterChangeRequestPeriodsOne& item) noexcept;

bool AbCounterChangeRequestPeriodsOneWriteJSON(std::ostream& s, const ::tl2::ab::CounterChangeRequestPeriodsOne& item) noexcept;
bool AbCounterChangeRequestPeriodsOneRead(::basictl::tl_istream & s, ::tl2::ab::CounterChangeRequestPeriodsOne& item) noexcept; 
bool AbCounterChangeRequestPeriodsOneWrite(::basictl::tl_ostream & s, const ::tl2::ab::CounterChangeRequestPeriodsOne& item) noexcept;
bool AbCounterChangeRequestPeriodsOneReadBoxed(::basictl::tl_istream & s, ::tl2::ab::CounterChangeRequestPeriodsOne& item);
bool AbCounterChangeRequestPeriodsOneWriteBoxed(::basictl::tl_ostream & s, const ::tl2::ab::CounterChangeRequestPeriodsOne& item);

}} // namespace tl2::details

