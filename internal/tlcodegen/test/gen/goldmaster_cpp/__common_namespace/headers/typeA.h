// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/typeA.h"

namespace tl2 { namespace details { 

void TypeAReset(::tl2::TypeA& item) noexcept;

bool TypeAWriteJSON(std::ostream& s, const ::tl2::TypeA& item) noexcept;
bool TypeARead(::basictl::tl_istream & s, ::tl2::TypeA& item) noexcept; 
bool TypeAWrite(::basictl::tl_ostream & s, const ::tl2::TypeA& item) noexcept;
bool TypeAReadBoxed(::basictl::tl_istream & s, ::tl2::TypeA& item);
bool TypeAWriteBoxed(::basictl::tl_ostream & s, const ::tl2::TypeA& item);

}} // namespace tl2::details

