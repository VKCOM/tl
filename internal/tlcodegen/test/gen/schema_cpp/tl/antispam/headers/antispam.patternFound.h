#pragma once

#include "../../../basics/basictl.h"
#include "../types/antispam.patternFound.h"

namespace tl2 { namespace details { 

void AntispamPatternFoundReset(::tl2::antispam::PatternFound& item);

bool AntispamPatternFoundWriteJSON(std::ostream& s, const ::tl2::antispam::PatternFound& item);
bool AntispamPatternFoundRead(::basictl::tl_istream & s, ::tl2::antispam::PatternFound& item);
bool AntispamPatternFoundWrite(::basictl::tl_ostream & s, const ::tl2::antispam::PatternFound& item);
bool AntispamPatternFoundReadBoxed(::basictl::tl_istream & s, ::tl2::antispam::PatternFound& item);
bool AntispamPatternFoundWriteBoxed(::basictl::tl_ostream & s, const ::tl2::antispam::PatternFound& item);

}} // namespace tl2::details

