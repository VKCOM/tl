// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/replace17.h"

namespace tlgen { namespace details { 

void Replace17Reset(::tlgen::Replace17& item) noexcept;

bool Replace17WriteJSON(std::ostream& s, const ::tlgen::Replace17& item) noexcept;
bool Replace17Read(::tlgen::basictl::tl_istream & s, ::tlgen::Replace17& item) noexcept; 
bool Replace17Write(::tlgen::basictl::tl_ostream & s, const ::tlgen::Replace17& item) noexcept;
bool Replace17ReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::Replace17& item);
bool Replace17WriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::Replace17& item);

}} // namespace tlgen::details

