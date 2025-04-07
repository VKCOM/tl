#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/service1.strvalueWithTime.h"

namespace tl2 { namespace details { 

void Service1StrvalueWithTimeReset(::tl2::service1::StrvalueWithTime& item);

bool Service1StrvalueWithTimeWriteJSON(std::ostream& s, const ::tl2::service1::StrvalueWithTime& item);
bool Service1StrvalueWithTimeRead(::basictl::tl_istream & s, ::tl2::service1::StrvalueWithTime& item);
bool Service1StrvalueWithTimeWrite(::basictl::tl_ostream & s, const ::tl2::service1::StrvalueWithTime& item);
bool Service1StrvalueWithTimeReadBoxed(::basictl::tl_istream & s, ::tl2::service1::StrvalueWithTime& item);
bool Service1StrvalueWithTimeWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::StrvalueWithTime& item);

}} // namespace tl2::details

