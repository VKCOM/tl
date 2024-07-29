#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/antispam.PatternFull.hpp"

namespace tl2 { namespace details { 

void AntispamPatternFullReset(::tl2::antispam::PatternFull& item);
bool AntispamPatternFullReadBoxed(::basictl::tl_istream & s, ::tl2::antispam::PatternFull& item);
bool AntispamPatternFullWriteBoxed(::basictl::tl_ostream & s, const ::tl2::antispam::PatternFull& item);

}} // namespace tl2::details

