// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/noStr.h"

namespace tlgen { namespace details { 

void NoStrReset(::tlgen::NoStr& item) noexcept;

bool NoStrWriteJSON(std::ostream& s, const ::tlgen::NoStr& item) noexcept;
bool NoStrRead(::tlgen::basictl::tl_istream & s, ::tlgen::NoStr& item) noexcept; 
bool NoStrWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::NoStr& item) noexcept;
bool NoStrReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::NoStr& item);
bool NoStrWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::NoStr& item);

}} // namespace tlgen::details

