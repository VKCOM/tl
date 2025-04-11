#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/antispam.patternNotFound.h"

namespace tl2 { namespace details { 

void AntispamPatternNotFoundReset(::tl2::antispam::PatternNotFound& item) noexcept;

bool AntispamPatternNotFoundWriteJSON(std::ostream& s, const ::tl2::antispam::PatternNotFound& item) noexcept;
bool AntispamPatternNotFoundRead(::basictl::tl_istream & s, ::tl2::antispam::PatternNotFound& item) noexcept; 
bool AntispamPatternNotFoundWrite(::basictl::tl_ostream & s, const ::tl2::antispam::PatternNotFound& item) noexcept;
bool AntispamPatternNotFoundReadBoxed(::basictl::tl_istream & s, ::tl2::antispam::PatternNotFound& item);
bool AntispamPatternNotFoundWriteBoxed(::basictl::tl_ostream & s, const ::tl2::antispam::PatternNotFound& item);

}} // namespace tl2::details

