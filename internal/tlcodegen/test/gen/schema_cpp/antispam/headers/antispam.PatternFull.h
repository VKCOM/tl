// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "antispam/types/antispam.PatternFull.h"

namespace tlgen { namespace details { 

void AntispamPatternFullReset(::tlgen::antispam::PatternFull& item) noexcept;

bool AntispamPatternFullWriteJSON(std::ostream & s, const ::tlgen::antispam::PatternFull& item) noexcept;
bool AntispamPatternFullReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::antispam::PatternFull& item) noexcept;
bool AntispamPatternFullWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::antispam::PatternFull& item) noexcept;

}} // namespace tlgen::details

