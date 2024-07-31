#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/antispam.getPattern.hpp"
#include "../types/antispam.PatternFull.hpp"

namespace tl2 { namespace details { 

void AntispamGetPatternReset(::tl2::antispam::GetPattern& item);

bool AntispamGetPatternWriteJSON(std::ostream& s, const ::tl2::antispam::GetPattern& item);
bool AntispamGetPatternRead(::basictl::tl_istream & s, ::tl2::antispam::GetPattern& item);
bool AntispamGetPatternWrite(::basictl::tl_ostream & s, const ::tl2::antispam::GetPattern& item);
bool AntispamGetPatternReadBoxed(::basictl::tl_istream & s, ::tl2::antispam::GetPattern& item);
bool AntispamGetPatternWriteBoxed(::basictl::tl_ostream & s, const ::tl2::antispam::GetPattern& item);

bool AntispamGetPatternReadResult(::basictl::tl_istream & s, ::tl2::antispam::GetPattern& item, ::tl2::antispam::PatternFull& result);
bool AntispamGetPatternWriteResult(::basictl::tl_ostream & s, ::tl2::antispam::GetPattern& item, ::tl2::antispam::PatternFull& result);
		
}} // namespace tl2::details

