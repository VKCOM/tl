// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "ab/types/ab.typeB.h"

namespace tl2 { namespace details { 

void AbTypeBReset(::tl2::ab::TypeB& item) noexcept;

bool AbTypeBWriteJSON(std::ostream& s, const ::tl2::ab::TypeB& item) noexcept;
bool AbTypeBRead(::basictl::tl_istream & s, ::tl2::ab::TypeB& item) noexcept; 
bool AbTypeBWrite(::basictl::tl_ostream & s, const ::tl2::ab::TypeB& item) noexcept;
bool AbTypeBReadBoxed(::basictl::tl_istream & s, ::tl2::ab::TypeB& item);
bool AbTypeBWriteBoxed(::basictl::tl_ostream & s, const ::tl2::ab::TypeB& item);

}} // namespace tl2::details

