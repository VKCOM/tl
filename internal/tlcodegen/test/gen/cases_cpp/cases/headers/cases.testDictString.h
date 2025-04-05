#pragma once

#include "../../basictl/io_streams.h"
#include "../types/cases.testDictString.h"

namespace tl2 { namespace details { 

void CasesTestDictStringReset(::tl2::cases::TestDictString& item);

bool CasesTestDictStringWriteJSON(std::ostream& s, const ::tl2::cases::TestDictString& item);
bool CasesTestDictStringRead(::basictl::tl_istream & s, ::tl2::cases::TestDictString& item);
bool CasesTestDictStringWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestDictString& item);
bool CasesTestDictStringReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestDictString& item);
bool CasesTestDictStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestDictString& item);

}} // namespace tl2::details

