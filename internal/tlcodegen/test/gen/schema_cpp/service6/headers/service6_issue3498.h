// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/issue3498.h"

namespace tlgen { namespace details { 

void Issue3498Reset(::tlgen::Issue3498& item) noexcept;

bool Issue3498WriteJSON(std::ostream& s, const ::tlgen::Issue3498& item) noexcept;
bool Issue3498Read(::tlgen::basictl::tl_istream & s, ::tlgen::Issue3498& item) noexcept; 
bool Issue3498Write(::tlgen::basictl::tl_ostream & s, const ::tlgen::Issue3498& item) noexcept;
bool Issue3498ReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::Issue3498& item);
bool Issue3498WriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::Issue3498& item);

}} // namespace tlgen::details

