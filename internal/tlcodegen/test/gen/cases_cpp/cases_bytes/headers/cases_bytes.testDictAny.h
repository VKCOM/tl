#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/cases_bytes.testDictAny.h"

namespace tl2 { namespace details { 

void CasesBytesTestDictAnyReset(::tl2::cases_bytes::TestDictAny& item) noexcept;

bool CasesBytesTestDictAnyWriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestDictAny& item) noexcept;
bool CasesBytesTestDictAnyRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictAny& item) noexcept; 
bool CasesBytesTestDictAnyWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictAny& item) noexcept;
bool CasesBytesTestDictAnyReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictAny& item);
bool CasesBytesTestDictAnyWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictAny& item);

}} // namespace tl2::details

