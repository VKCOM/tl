#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/cases_bytes.testDictStringString.h"

namespace tl2 { namespace details { 

void CasesBytesTestDictStringStringReset(::tl2::cases_bytes::TestDictStringString& item) noexcept;

bool CasesBytesTestDictStringStringWriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestDictStringString& item) noexcept;
bool CasesBytesTestDictStringStringRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictStringString& item) noexcept; 
bool CasesBytesTestDictStringStringWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictStringString& item) noexcept;
bool CasesBytesTestDictStringStringReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictStringString& item);
bool CasesBytesTestDictStringStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictStringString& item);

}} // namespace tl2::details

