// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/int32.h"

namespace tlgen { namespace details { 

void Int32Reset(::tlgen::Int32& item) noexcept;

bool Int32WriteJSON(std::ostream& s, const ::tlgen::Int32& item) noexcept;
bool Int32Read(::tlgen::basictl::tl_istream & s, ::tlgen::Int32& item) noexcept; 
bool Int32Write(::tlgen::basictl::tl_ostream & s, const ::tlgen::Int32& item) noexcept;
bool Int32ReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::Int32& item);
bool Int32WriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::Int32& item);

}} // namespace tlgen::details

