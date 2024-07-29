#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/antispam.patternFound.hpp"

namespace tl2 { namespace details { 

void AntispamPatternFoundReset(::tl2::antispam::PatternFound& item);
bool AntispamPatternFoundRead(::basictl::tl_istream & s, ::tl2::antispam::PatternFound& item);
bool AntispamPatternFoundWrite(::basictl::tl_ostream & s, const ::tl2::antispam::PatternFound& item);
bool AntispamPatternFoundReadBoxed(::basictl::tl_istream & s, ::tl2::antispam::PatternFound& item);
bool AntispamPatternFoundWriteBoxed(::basictl::tl_ostream & s, const ::tl2::antispam::PatternFound& item);

}} // namespace tl2::details

