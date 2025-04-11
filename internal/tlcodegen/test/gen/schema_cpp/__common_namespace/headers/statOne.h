#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/statOne.h"

namespace tl2 { namespace details { 

void StatOneReset(::tl2::StatOne& item) noexcept;

bool StatOneWriteJSON(std::ostream& s, const ::tl2::StatOne& item) noexcept;
bool StatOneRead(::basictl::tl_istream & s, ::tl2::StatOne& item) noexcept; 
bool StatOneWrite(::basictl::tl_ostream & s, const ::tl2::StatOne& item) noexcept;
bool StatOneReadBoxed(::basictl::tl_istream & s, ::tl2::StatOne& item);
bool StatOneWriteBoxed(::basictl::tl_ostream & s, const ::tl2::StatOne& item);

}} // namespace tl2::details

