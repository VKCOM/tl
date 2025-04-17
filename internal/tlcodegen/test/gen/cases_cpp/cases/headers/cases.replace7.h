#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases/types/cases.replace7.h"

namespace tl2 { namespace details { 

void CasesReplace7Reset(::tl2::cases::Replace7& item) noexcept;

bool CasesReplace7WriteJSON(std::ostream& s, const ::tl2::cases::Replace7& item) noexcept;
bool CasesReplace7Read(::basictl::tl_istream & s, ::tl2::cases::Replace7& item) noexcept; 
bool CasesReplace7Write(::basictl::tl_ostream & s, const ::tl2::cases::Replace7& item) noexcept;
bool CasesReplace7ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::Replace7& item);
bool CasesReplace7WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::Replace7& item);

}} // namespace tl2::details

