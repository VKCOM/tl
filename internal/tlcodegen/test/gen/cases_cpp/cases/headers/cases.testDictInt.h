#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/cases.testDictInt.h"

namespace tl2 { namespace details { 

void CasesTestDictIntReset(::tl2::cases::TestDictInt& item) noexcept;

bool CasesTestDictIntWriteJSON(std::ostream& s, const ::tl2::cases::TestDictInt& item) noexcept;
bool CasesTestDictIntRead(::basictl::tl_istream & s, ::tl2::cases::TestDictInt& item) noexcept; 
bool CasesTestDictIntWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestDictInt& item) noexcept;
bool CasesTestDictIntReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestDictInt& item);
bool CasesTestDictIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestDictInt& item);

}} // namespace tl2::details

