// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/typeB.h"

namespace tl2 { namespace details { 

void TypeBReset(::tl2::TypeB& item) noexcept;

bool TypeBWriteJSON(std::ostream& s, const ::tl2::TypeB& item) noexcept;
bool TypeBRead(::basictl::tl_istream & s, ::tl2::TypeB& item) noexcept; 
bool TypeBWrite(::basictl::tl_ostream & s, const ::tl2::TypeB& item) noexcept;
bool TypeBReadBoxed(::basictl::tl_istream & s, ::tl2::TypeB& item);
bool TypeBWriteBoxed(::basictl::tl_ostream & s, const ::tl2::TypeB& item);

}} // namespace tl2::details

