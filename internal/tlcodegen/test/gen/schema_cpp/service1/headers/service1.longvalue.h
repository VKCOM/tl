// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service1/types/service1.longvalue.h"

namespace tl2 { namespace details { 

void Service1LongvalueReset(::tl2::service1::Longvalue& item) noexcept;

bool Service1LongvalueWriteJSON(std::ostream& s, const ::tl2::service1::Longvalue& item) noexcept;
bool Service1LongvalueRead(::basictl::tl_istream & s, ::tl2::service1::Longvalue& item) noexcept; 
bool Service1LongvalueWrite(::basictl::tl_ostream & s, const ::tl2::service1::Longvalue& item) noexcept;
bool Service1LongvalueReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Longvalue& item);
bool Service1LongvalueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Longvalue& item);

}} // namespace tl2::details

