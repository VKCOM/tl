#pragma once

#include "../../../basics/basictl.h"
#include "../types/cases.replace7plus.h"

namespace tl2 { namespace details { 

void CasesReplace7plusReset(::tl2::cases::Replace7plus& item);

bool CasesReplace7plusWriteJSON(std::ostream& s, const ::tl2::cases::Replace7plus& item);
bool CasesReplace7plusRead(::basictl::tl_istream & s, ::tl2::cases::Replace7plus& item);
bool CasesReplace7plusWrite(::basictl::tl_ostream & s, const ::tl2::cases::Replace7plus& item);
bool CasesReplace7plusReadBoxed(::basictl::tl_istream & s, ::tl2::cases::Replace7plus& item);
bool CasesReplace7plusWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::Replace7plus& item);

}} // namespace tl2::details

