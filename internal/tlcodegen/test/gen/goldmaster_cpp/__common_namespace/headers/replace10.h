// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/replace10.h"

namespace tlgen { namespace details { 

void Replace10Reset(::tlgen::Replace10& item) noexcept;

bool Replace10WriteJSON(std::ostream& s, const ::tlgen::Replace10& item) noexcept;
bool Replace10Read(::tlgen::basictl::tl_istream & s, ::tlgen::Replace10& item) noexcept; 
bool Replace10Write(::tlgen::basictl::tl_ostream & s, const ::tlgen::Replace10& item) noexcept;
bool Replace10ReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::Replace10& item);
bool Replace10WriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::Replace10& item);

}} // namespace tlgen::details

