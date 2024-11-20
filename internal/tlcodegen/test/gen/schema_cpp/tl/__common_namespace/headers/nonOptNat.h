#pragma once

#include "../../../basics/basictl.h"
#include "../types/nonOptNat.h"

namespace tl2 { namespace details { 

void NonOptNatReset(::tl2::NonOptNat& item);

bool NonOptNatWriteJSON(std::ostream& s, const ::tl2::NonOptNat& item);
bool NonOptNatRead(::basictl::tl_istream & s, ::tl2::NonOptNat& item);
bool NonOptNatWrite(::basictl::tl_ostream & s, const ::tl2::NonOptNat& item);
bool NonOptNatReadBoxed(::basictl::tl_istream & s, ::tl2::NonOptNat& item);
bool NonOptNatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::NonOptNat& item);

}} // namespace tl2::details

