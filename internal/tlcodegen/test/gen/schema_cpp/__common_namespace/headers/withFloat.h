// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/withFloat.h"

namespace tl2 { namespace details { 

void WithFloatReset(::tl2::WithFloat& item) noexcept;

bool WithFloatWriteJSON(std::ostream& s, const ::tl2::WithFloat& item) noexcept;
bool WithFloatRead(::basictl::tl_istream & s, ::tl2::WithFloat& item) noexcept; 
bool WithFloatWrite(::basictl::tl_ostream & s, const ::tl2::WithFloat& item) noexcept;
bool WithFloatReadBoxed(::basictl::tl_istream & s, ::tl2::WithFloat& item);
bool WithFloatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::WithFloat& item);

}} // namespace tl2::details

