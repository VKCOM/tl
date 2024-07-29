#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/cases.testLocalFieldmask.hpp"

namespace tl2 { namespace details { 

void CasesTestLocalFieldmaskReset(::tl2::cases::TestLocalFieldmask& item);
bool CasesTestLocalFieldmaskRead(::basictl::tl_istream & s, ::tl2::cases::TestLocalFieldmask& item);
bool CasesTestLocalFieldmaskWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestLocalFieldmask& item);
bool CasesTestLocalFieldmaskReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestLocalFieldmask& item);
bool CasesTestLocalFieldmaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestLocalFieldmask& item);

}} // namespace tl2::details

