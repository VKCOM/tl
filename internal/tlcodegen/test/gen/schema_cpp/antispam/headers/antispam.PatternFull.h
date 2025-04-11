#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/antispam.PatternFull.h"

namespace tl2 { namespace details { 

void AntispamPatternFullReset(::tl2::antispam::PatternFull& item) noexcept;

bool AntispamPatternFullWriteJSON(std::ostream & s, const ::tl2::antispam::PatternFull& item) noexcept;
bool AntispamPatternFullReadBoxed(::basictl::tl_istream & s, ::tl2::antispam::PatternFull& item) noexcept;
bool AntispamPatternFullWriteBoxed(::basictl::tl_ostream & s, const ::tl2::antispam::PatternFull& item) noexcept;

}} // namespace tl2::details

