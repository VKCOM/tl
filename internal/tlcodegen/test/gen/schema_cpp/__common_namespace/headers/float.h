// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/float.h"

namespace tl2 { namespace details { 

void FloatReset(float& item) noexcept;

bool FloatWriteJSON(std::ostream& s, const float& item) noexcept;
bool FloatRead(::basictl::tl_istream & s, float& item) noexcept; 
bool FloatWrite(::basictl::tl_ostream & s, const float& item) noexcept;
bool FloatReadBoxed(::basictl::tl_istream & s, float& item);
bool FloatWriteBoxed(::basictl::tl_ostream & s, const float& item);

}} // namespace tl2::details

