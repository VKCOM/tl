#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service1.strvalue.hpp"

namespace tl2 { namespace details { 

void Service1StrvalueReset(::tl2::service1::Strvalue& item);
bool Service1StrvalueRead(::basictl::tl_istream & s, ::tl2::service1::Strvalue& item);
bool Service1StrvalueWrite(::basictl::tl_ostream & s, const ::tl2::service1::Strvalue& item);
bool Service1StrvalueReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Strvalue& item);
bool Service1StrvalueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Strvalue& item);

}} // namespace tl2::details

