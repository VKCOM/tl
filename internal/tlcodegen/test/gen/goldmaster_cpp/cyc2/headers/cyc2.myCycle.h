// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cyc2/types/cyc2.myCycle.h"

namespace tl2 { namespace details { 

void Cyc2MyCycleReset(::tl2::cyc2::MyCycle& item) noexcept;

bool Cyc2MyCycleWriteJSON(std::ostream& s, const ::tl2::cyc2::MyCycle& item) noexcept;
bool Cyc2MyCycleRead(::basictl::tl_istream & s, ::tl2::cyc2::MyCycle& item) noexcept; 
bool Cyc2MyCycleWrite(::basictl::tl_ostream & s, const ::tl2::cyc2::MyCycle& item) noexcept;
bool Cyc2MyCycleReadBoxed(::basictl::tl_istream & s, ::tl2::cyc2::MyCycle& item);
bool Cyc2MyCycleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cyc2::MyCycle& item);

}} // namespace tl2::details

