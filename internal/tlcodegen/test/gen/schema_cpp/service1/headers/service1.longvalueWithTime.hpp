#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service1.longvalueWithTime.hpp"

namespace tl2 { namespace details { 

void Service1LongvalueWithTimeReset(::tl2::service1::LongvalueWithTime& item);

bool Service1LongvalueWithTimeWriteJSON(std::ostream& s, const ::tl2::service1::LongvalueWithTime& item);
bool Service1LongvalueWithTimeRead(::basictl::tl_istream & s, ::tl2::service1::LongvalueWithTime& item);
bool Service1LongvalueWithTimeWrite(::basictl::tl_ostream & s, const ::tl2::service1::LongvalueWithTime& item);
bool Service1LongvalueWithTimeReadBoxed(::basictl::tl_istream & s, ::tl2::service1::LongvalueWithTime& item);
bool Service1LongvalueWithTimeWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::LongvalueWithTime& item);

}} // namespace tl2::details

