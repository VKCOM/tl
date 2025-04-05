#pragma once

#include "../../basictl/io_streams.h"
#include "../types/antispam.patternNotFound.h"

namespace tl2 { namespace details { 

void AntispamPatternNotFoundReset(::tl2::antispam::PatternNotFound& item);

bool AntispamPatternNotFoundWriteJSON(std::ostream& s, const ::tl2::antispam::PatternNotFound& item);
bool AntispamPatternNotFoundRead(::basictl::tl_istream & s, ::tl2::antispam::PatternNotFound& item);
bool AntispamPatternNotFoundWrite(::basictl::tl_ostream & s, const ::tl2::antispam::PatternNotFound& item);
bool AntispamPatternNotFoundReadBoxed(::basictl::tl_istream & s, ::tl2::antispam::PatternNotFound& item);
bool AntispamPatternNotFoundWriteBoxed(::basictl::tl_ostream & s, const ::tl2::antispam::PatternNotFound& item);

}} // namespace tl2::details

