#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases_bytes/types/cases_bytes.testDictString.h"

namespace tl2 { namespace details { 

void CasesBytesTestDictStringReset(::tl2::cases_bytes::TestDictString& item) noexcept;

bool CasesBytesTestDictStringWriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestDictString& item) noexcept;
bool CasesBytesTestDictStringRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictString& item) noexcept; 
bool CasesBytesTestDictStringWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictString& item) noexcept;
bool CasesBytesTestDictStringReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictString& item);
bool CasesBytesTestDictStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictString& item);

}} // namespace tl2::details

