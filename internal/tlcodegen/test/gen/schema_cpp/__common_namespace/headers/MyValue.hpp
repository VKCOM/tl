#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/MyValue.hpp"

namespace tl2 { namespace details { 

void MyValueReset(::tl2::MyValue& item);
bool MyValueReadBoxed(::basictl::tl_istream & s, ::tl2::MyValue& item);
bool MyValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyValue& item);

}} // namespace tl2::details

