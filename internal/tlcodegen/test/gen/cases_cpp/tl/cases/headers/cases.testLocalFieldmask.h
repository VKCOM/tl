#pragma once

#include "../../../basics/basictl.h"
#include "../types/cases.testLocalFieldmask.h"

namespace tl2 { namespace details { 

void CasesTestLocalFieldmaskReset(::tl2::cases::TestLocalFieldmask& item);

bool CasesTestLocalFieldmaskWriteJSON(std::ostream& s, const ::tl2::cases::TestLocalFieldmask& item);
bool CasesTestLocalFieldmaskRead(::basictl::tl_istream & s, ::tl2::cases::TestLocalFieldmask& item);
bool CasesTestLocalFieldmaskWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestLocalFieldmask& item);
bool CasesTestLocalFieldmaskReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestLocalFieldmask& item);
bool CasesTestLocalFieldmaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestLocalFieldmask& item);

}} // namespace tl2::details

