#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../types/service1.strvalueWithTime.hpp"

namespace tl2 { namespace details { 

void Service1StrvalueWithTimeReset(::tl2::service1::StrvalueWithTime& item);
bool Service1StrvalueWithTimeRead(::basictl::tl_istream & s, ::tl2::service1::StrvalueWithTime& item);
bool Service1StrvalueWithTimeWrite(::basictl::tl_ostream & s, const ::tl2::service1::StrvalueWithTime& item);
bool Service1StrvalueWithTimeReadBoxed(::basictl::tl_istream & s, ::tl2::service1::StrvalueWithTime& item);
bool Service1StrvalueWithTimeWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::StrvalueWithTime& item);

}} // namespace tl2::details

