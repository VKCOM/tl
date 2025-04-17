#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service1/types/service1.strvalue.h"

namespace tl2 { namespace details { 

void Service1StrvalueReset(::tl2::service1::Strvalue& item) noexcept;

bool Service1StrvalueWriteJSON(std::ostream& s, const ::tl2::service1::Strvalue& item) noexcept;
bool Service1StrvalueRead(::basictl::tl_istream & s, ::tl2::service1::Strvalue& item) noexcept; 
bool Service1StrvalueWrite(::basictl::tl_ostream & s, const ::tl2::service1::Strvalue& item) noexcept;
bool Service1StrvalueReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Strvalue& item);
bool Service1StrvalueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Strvalue& item);

}} // namespace tl2::details

