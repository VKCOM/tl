#pragma once

#include "../../../basics/basictl.h"
#include "../types/service1.longvalue.h"

namespace tl2 { namespace details { 

void Service1LongvalueReset(::tl2::service1::Longvalue& item);

bool Service1LongvalueWriteJSON(std::ostream& s, const ::tl2::service1::Longvalue& item);
bool Service1LongvalueRead(::basictl::tl_istream & s, ::tl2::service1::Longvalue& item);
bool Service1LongvalueWrite(::basictl::tl_ostream & s, const ::tl2::service1::Longvalue& item);
bool Service1LongvalueReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Longvalue& item);
bool Service1LongvalueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Longvalue& item);

}} // namespace tl2::details

