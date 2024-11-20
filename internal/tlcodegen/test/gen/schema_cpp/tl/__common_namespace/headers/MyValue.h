#pragma once

#include "../../../basics/basictl.h"
#include "../types/MyValue.h"

namespace tl2 { namespace details { 

void MyValueReset(::tl2::MyValue& item);

bool MyValueWriteJSON(std::ostream & s, const ::tl2::MyValue& item);
bool MyValueReadBoxed(::basictl::tl_istream & s, ::tl2::MyValue& item);
bool MyValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyValue& item);

}} // namespace tl2::details

