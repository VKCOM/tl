#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/cases_bytes.testDictStringString.hpp"

namespace tl2 { namespace details { 

void CasesBytesTestDictStringStringReset(::tl2::cases_bytes::TestDictStringString& item);
bool CasesBytesTestDictStringStringRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictStringString& item);
bool CasesBytesTestDictStringStringWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictStringString& item);
bool CasesBytesTestDictStringStringReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictStringString& item);
bool CasesBytesTestDictStringStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictStringString& item);

}} // namespace tl2::details

