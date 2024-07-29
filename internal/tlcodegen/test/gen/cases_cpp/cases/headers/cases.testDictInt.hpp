#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/cases.testDictInt.hpp"

namespace tl2 { namespace details { 

void CasesTestDictIntReset(::tl2::cases::TestDictInt& item);
bool CasesTestDictIntRead(::basictl::tl_istream & s, ::tl2::cases::TestDictInt& item);
bool CasesTestDictIntWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestDictInt& item);
bool CasesTestDictIntReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestDictInt& item);
bool CasesTestDictIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestDictInt& item);

}} // namespace tl2::details

