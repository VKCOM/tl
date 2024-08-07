#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/cases_bytes.testTuple.hpp"

namespace tl2 { namespace details { 

void CasesBytesTestTupleReset(::tl2::cases_bytes::TestTuple& item);

bool CasesBytesTestTupleWriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestTuple& item);
bool CasesBytesTestTupleRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestTuple& item);
bool CasesBytesTestTupleWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestTuple& item);
bool CasesBytesTestTupleReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestTuple& item);
bool CasesBytesTestTupleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestTuple& item);

}} // namespace tl2::details

