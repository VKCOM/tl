// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/vector.h"
#include "cyc1/types/cyc1.myCycle.h"

namespace tl2 { namespace details { 

void VectorCyc1MyCycleReset(std::vector<::tl2::cyc1::MyCycle>& item) noexcept;

bool VectorCyc1MyCycleWriteJSON(std::ostream& s, const std::vector<::tl2::cyc1::MyCycle>& item) noexcept;
bool VectorCyc1MyCycleRead(::basictl::tl_istream & s, std::vector<::tl2::cyc1::MyCycle>& item) noexcept; 
bool VectorCyc1MyCycleWrite(::basictl::tl_ostream & s, const std::vector<::tl2::cyc1::MyCycle>& item) noexcept;
bool VectorCyc1MyCycleReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::cyc1::MyCycle>& item);
bool VectorCyc1MyCycleWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::cyc1::MyCycle>& item);

}} // namespace tl2::details

