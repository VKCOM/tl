// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "ab/types/ab.typeA.h"

namespace tlgen { namespace details { 

void AbTypeAReset(::tlgen::ab::TypeA& item) noexcept;

bool AbTypeAWriteJSON(std::ostream& s, const ::tlgen::ab::TypeA& item) noexcept;
bool AbTypeARead(::tlgen::basictl::tl_istream & s, ::tlgen::ab::TypeA& item) noexcept; 
bool AbTypeAWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::ab::TypeA& item) noexcept;
bool AbTypeAReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::ab::TypeA& item);
bool AbTypeAWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::ab::TypeA& item);

}} // namespace tlgen::details

