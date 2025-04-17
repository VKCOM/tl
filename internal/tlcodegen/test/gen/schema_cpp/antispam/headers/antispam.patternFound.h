#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "antispam/types/antispam.patternFound.h"

namespace tl2 { namespace details { 

void AntispamPatternFoundReset(::tl2::antispam::PatternFound& item) noexcept;

bool AntispamPatternFoundWriteJSON(std::ostream& s, const ::tl2::antispam::PatternFound& item) noexcept;
bool AntispamPatternFoundRead(::basictl::tl_istream & s, ::tl2::antispam::PatternFound& item) noexcept; 
bool AntispamPatternFoundWrite(::basictl::tl_ostream & s, const ::tl2::antispam::PatternFound& item) noexcept;
bool AntispamPatternFoundReadBoxed(::basictl::tl_istream & s, ::tl2::antispam::PatternFound& item);
bool AntispamPatternFoundWriteBoxed(::basictl::tl_ostream & s, const ::tl2::antispam::PatternFound& item);

}} // namespace tl2::details

