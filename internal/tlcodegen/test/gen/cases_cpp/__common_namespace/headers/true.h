#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/true.h"

namespace tl2 { namespace details { 

void TrueReset(::tl2::True& item) noexcept;

bool TrueWriteJSON(std::ostream& s, const ::tl2::True& item) noexcept;
bool TrueRead(::basictl::tl_istream & s, ::tl2::True& item) noexcept; 
bool TrueWrite(::basictl::tl_ostream & s, const ::tl2::True& item) noexcept;
bool TrueReadBoxed(::basictl::tl_istream & s, ::tl2::True& item);
bool TrueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::True& item);

}} // namespace tl2::details

