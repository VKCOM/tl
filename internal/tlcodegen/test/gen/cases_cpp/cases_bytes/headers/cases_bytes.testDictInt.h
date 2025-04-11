#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/cases_bytes.testDictInt.h"

namespace tl2 { namespace details { 

void CasesBytesTestDictIntReset(::tl2::cases_bytes::TestDictInt& item) noexcept;

bool CasesBytesTestDictIntWriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestDictInt& item) noexcept;
bool CasesBytesTestDictIntRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictInt& item) noexcept; 
bool CasesBytesTestDictIntWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictInt& item) noexcept;
bool CasesBytesTestDictIntReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictInt& item);
bool CasesBytesTestDictIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictInt& item);

}} // namespace tl2::details

