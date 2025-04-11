#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/cases.testTuple.h"

namespace tl2 { namespace details { 

void CasesTestTupleReset(::tl2::cases::TestTuple& item) noexcept;

bool CasesTestTupleWriteJSON(std::ostream& s, const ::tl2::cases::TestTuple& item) noexcept;
bool CasesTestTupleRead(::basictl::tl_istream & s, ::tl2::cases::TestTuple& item) noexcept; 
bool CasesTestTupleWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestTuple& item) noexcept;
bool CasesTestTupleReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestTuple& item);
bool CasesTestTupleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestTuple& item);

}} // namespace tl2::details

