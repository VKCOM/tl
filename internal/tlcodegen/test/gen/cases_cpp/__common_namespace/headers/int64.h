// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/int64.h"

namespace tlgen { namespace details { 

void Int64Reset(::tlgen::Int64& item) noexcept;

bool Int64WriteJSON(std::ostream& s, const ::tlgen::Int64& item) noexcept;
bool Int64Read(::tlgen::basictl::tl_istream & s, ::tlgen::Int64& item) noexcept; 
bool Int64Write(::tlgen::basictl::tl_ostream & s, const ::tlgen::Int64& item) noexcept;
bool Int64ReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::Int64& item);
bool Int64WriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::Int64& item);

}} // namespace tlgen::details

