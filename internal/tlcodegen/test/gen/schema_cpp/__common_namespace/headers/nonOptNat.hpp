#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/nonOptNat.hpp"

namespace tl2 { namespace details { 

void NonOptNatReset(::tl2::NonOptNat& item);
bool NonOptNatRead(::basictl::tl_istream & s, ::tl2::NonOptNat& item);
bool NonOptNatWrite(::basictl::tl_ostream & s, const ::tl2::NonOptNat& item);
bool NonOptNatReadBoxed(::basictl::tl_istream & s, ::tl2::NonOptNat& item);
bool NonOptNatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::NonOptNat& item);

}} // namespace tl2::details

