// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "antispam/functions/antispam.getPattern.h"
#include "antispam/types/antispam.PatternFull.h"

namespace tl2 { namespace details { 

void AntispamGetPatternReset(::tl2::antispam::GetPattern& item) noexcept;

bool AntispamGetPatternWriteJSON(std::ostream& s, const ::tl2::antispam::GetPattern& item) noexcept;
bool AntispamGetPatternRead(::basictl::tl_istream & s, ::tl2::antispam::GetPattern& item) noexcept; 
bool AntispamGetPatternWrite(::basictl::tl_ostream & s, const ::tl2::antispam::GetPattern& item) noexcept;
bool AntispamGetPatternReadBoxed(::basictl::tl_istream & s, ::tl2::antispam::GetPattern& item);
bool AntispamGetPatternWriteBoxed(::basictl::tl_ostream & s, const ::tl2::antispam::GetPattern& item);

bool AntispamGetPatternReadResult(::basictl::tl_istream & s, ::tl2::antispam::GetPattern& item, ::tl2::antispam::PatternFull& result);
bool AntispamGetPatternWriteResult(::basictl::tl_ostream & s, ::tl2::antispam::GetPattern& item, ::tl2::antispam::PatternFull& result);
		
}} // namespace tl2::details

