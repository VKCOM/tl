// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases/types/cases.testTuple.h"

namespace tlgen { namespace details { 

void CasesTestTupleReset(::tlgen::cases::TestTuple& item) noexcept;

bool CasesTestTupleWriteJSON(std::ostream& s, const ::tlgen::cases::TestTuple& item) noexcept;
bool CasesTestTupleRead(::tlgen::basictl::tl_istream & s, ::tlgen::cases::TestTuple& item) noexcept; 
bool CasesTestTupleWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::cases::TestTuple& item) noexcept;
bool CasesTestTupleReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::cases::TestTuple& item);
bool CasesTestTupleWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::cases::TestTuple& item);

}} // namespace tlgen::details

