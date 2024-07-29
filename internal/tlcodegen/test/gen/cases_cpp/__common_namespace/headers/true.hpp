#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/true.hpp"

namespace tl2 { namespace details { 

void TrueReset(::tl2::True& item);
bool TrueRead(::basictl::tl_istream & s, ::tl2::True& item);
bool TrueWrite(::basictl::tl_ostream & s, const ::tl2::True& item);
bool TrueReadBoxed(::basictl::tl_istream & s, ::tl2::True& item);
bool TrueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::True& item);

}} // namespace tl2::details

