#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/cases.testDictAny.hpp"

namespace tl2 { namespace details { 

void CasesTestDictAnyReset(::tl2::cases::TestDictAny& item);

bool CasesTestDictAnyWriteJSON(std::ostream& s, const ::tl2::cases::TestDictAny& item);
bool CasesTestDictAnyRead(::basictl::tl_istream & s, ::tl2::cases::TestDictAny& item);
bool CasesTestDictAnyWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestDictAny& item);
bool CasesTestDictAnyReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestDictAny& item);
bool CasesTestDictAnyWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestDictAny& item);

}} // namespace tl2::details

