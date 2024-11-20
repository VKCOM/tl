#pragma once

#include "../../../basics/basictl.h"
#include "../types/cases_bytes.testDictString.h"

namespace tl2 { namespace details { 

void CasesBytesTestDictStringReset(::tl2::cases_bytes::TestDictString& item);

bool CasesBytesTestDictStringWriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestDictString& item);
bool CasesBytesTestDictStringRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictString& item);
bool CasesBytesTestDictStringWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictString& item);
bool CasesBytesTestDictStringReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictString& item);
bool CasesBytesTestDictStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictString& item);

}} // namespace tl2::details

