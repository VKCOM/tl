// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "ab/types/ab.typeC.h"

namespace tl2 { namespace details { 

void AbTypeCReset(::tl2::ab::TypeC& item) noexcept;

bool AbTypeCWriteJSON(std::ostream& s, const ::tl2::ab::TypeC& item) noexcept;
bool AbTypeCRead(::basictl::tl_istream & s, ::tl2::ab::TypeC& item) noexcept; 
bool AbTypeCWrite(::basictl::tl_ostream & s, const ::tl2::ab::TypeC& item) noexcept;
bool AbTypeCReadBoxed(::basictl::tl_istream & s, ::tl2::ab::TypeC& item);
bool AbTypeCWriteBoxed(::basictl::tl_ostream & s, const ::tl2::ab::TypeC& item);

}} // namespace tl2::details

