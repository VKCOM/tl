#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/nonOptNat.h"

namespace tl2 { namespace details { 

void NonOptNatReset(::tl2::NonOptNat& item) noexcept;

bool NonOptNatWriteJSON(std::ostream& s, const ::tl2::NonOptNat& item) noexcept;
bool NonOptNatRead(::basictl::tl_istream & s, ::tl2::NonOptNat& item) noexcept; 
bool NonOptNatWrite(::basictl::tl_ostream & s, const ::tl2::NonOptNat& item) noexcept;
bool NonOptNatReadBoxed(::basictl::tl_istream & s, ::tl2::NonOptNat& item);
bool NonOptNatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::NonOptNat& item);

}} // namespace tl2::details

