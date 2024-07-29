#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/cases_bytes.testArray.hpp"

namespace tl2 { namespace details { 

void CasesBytesTestArrayReset(::tl2::cases_bytes::TestArray& item);
bool CasesBytesTestArrayRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestArray& item);
bool CasesBytesTestArrayWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestArray& item);
bool CasesBytesTestArrayReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestArray& item);
bool CasesBytesTestArrayWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestArray& item);

}} // namespace tl2::details

