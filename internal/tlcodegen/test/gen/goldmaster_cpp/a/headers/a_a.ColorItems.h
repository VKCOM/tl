// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "b/types/a.ColorItems.h"

namespace tl2 { namespace details { 

void BRedReset(::tl2::b::Red& item) noexcept;

bool BRedWriteJSON(std::ostream& s, const ::tl2::b::Red& item) noexcept;
bool BRedRead(::basictl::tl_istream & s, ::tl2::b::Red& item) noexcept; 
bool BRedWrite(::basictl::tl_ostream & s, const ::tl2::b::Red& item) noexcept;
bool BRedReadBoxed(::basictl::tl_istream & s, ::tl2::b::Red& item);
bool BRedWriteBoxed(::basictl::tl_ostream & s, const ::tl2::b::Red& item);

}} // namespace tl2::details

