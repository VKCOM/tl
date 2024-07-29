#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/withFloat.hpp"

namespace tl2 { namespace details { 

void WithFloatReset(::tl2::WithFloat& item);
bool WithFloatRead(::basictl::tl_istream & s, ::tl2::WithFloat& item);
bool WithFloatWrite(::basictl::tl_ostream & s, const ::tl2::WithFloat& item);
bool WithFloatReadBoxed(::basictl::tl_istream & s, ::tl2::WithFloat& item);
bool WithFloatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::WithFloat& item);

}} // namespace tl2::details

