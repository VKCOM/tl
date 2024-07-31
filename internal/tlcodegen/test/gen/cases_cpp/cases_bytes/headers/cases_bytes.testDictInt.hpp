#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/cases_bytes.testDictInt.hpp"

namespace tl2 { namespace details { 

void CasesBytesTestDictIntReset(::tl2::cases_bytes::TestDictInt& item);

bool CasesBytesTestDictIntWriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestDictInt& item);
bool CasesBytesTestDictIntRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictInt& item);
bool CasesBytesTestDictIntWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictInt& item);
bool CasesBytesTestDictIntReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictInt& item);
bool CasesBytesTestDictIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictInt& item);

}} // namespace tl2::details

