#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases_bytes/types/cases_bytes.testVector.h"

namespace tl2 { namespace details { 

void CasesBytesTestVectorReset(::tl2::cases_bytes::TestVector& item) noexcept;

bool CasesBytesTestVectorWriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestVector& item) noexcept;
bool CasesBytesTestVectorRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestVector& item) noexcept; 
bool CasesBytesTestVectorWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestVector& item) noexcept;
bool CasesBytesTestVectorReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestVector& item);
bool CasesBytesTestVectorWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestVector& item);

}} // namespace tl2::details

