#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service1.longvalue.hpp"

namespace tl2 { namespace details { 

void Service1LongvalueReset(::tl2::service1::Longvalue& item);
bool Service1LongvalueRead(::basictl::tl_istream & s, ::tl2::service1::Longvalue& item);
bool Service1LongvalueWrite(::basictl::tl_ostream & s, const ::tl2::service1::Longvalue& item);
bool Service1LongvalueReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Longvalue& item);
bool Service1LongvalueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Longvalue& item);

}} // namespace tl2::details

