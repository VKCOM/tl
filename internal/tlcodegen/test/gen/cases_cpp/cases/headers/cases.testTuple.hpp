#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/cases.testTuple.hpp"

namespace tl2 { namespace details { 

void CasesTestTupleReset(::tl2::cases::TestTuple& item);
bool CasesTestTupleRead(::basictl::tl_istream & s, ::tl2::cases::TestTuple& item);
bool CasesTestTupleWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestTuple& item);
bool CasesTestTupleReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestTuple& item);
bool CasesTestTupleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestTuple& item);

}} // namespace tl2::details

