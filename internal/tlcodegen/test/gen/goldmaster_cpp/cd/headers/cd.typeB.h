// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cd/types/cd.typeB.h"

namespace tlgen { namespace details { 

void CdTypeBReset(::tlgen::cd::TypeB& item) noexcept;

bool CdTypeBWriteJSON(std::ostream& s, const ::tlgen::cd::TypeB& item) noexcept;
bool CdTypeBRead(::tlgen::basictl::tl_istream & s, ::tlgen::cd::TypeB& item) noexcept; 
bool CdTypeBWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::cd::TypeB& item) noexcept;
bool CdTypeBReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::cd::TypeB& item);
bool CdTypeBWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::cd::TypeB& item);

}} // namespace tlgen::details

