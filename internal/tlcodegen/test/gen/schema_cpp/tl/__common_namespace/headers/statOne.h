#pragma once

#include "../../../basics/basictl.h"
#include "../types/statOne.h"

namespace tl2 { namespace details { 

void StatOneReset(::tl2::StatOne& item);

bool StatOneWriteJSON(std::ostream& s, const ::tl2::StatOne& item);
bool StatOneRead(::basictl::tl_istream & s, ::tl2::StatOne& item);
bool StatOneWrite(::basictl::tl_ostream & s, const ::tl2::StatOne& item);
bool StatOneReadBoxed(::basictl::tl_istream & s, ::tl2::StatOne& item);
bool StatOneWriteBoxed(::basictl::tl_ostream & s, const ::tl2::StatOne& item);

}} // namespace tl2::details

