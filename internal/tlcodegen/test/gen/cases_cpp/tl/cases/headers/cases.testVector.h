#pragma once

#include "../../../basics/basictl.h"
#include "../types/cases.testVector.h"

namespace tl2 { namespace details { 

void CasesTestVectorReset(::tl2::cases::TestVector& item);

bool CasesTestVectorWriteJSON(std::ostream& s, const ::tl2::cases::TestVector& item);
bool CasesTestVectorRead(::basictl::tl_istream & s, ::tl2::cases::TestVector& item);
bool CasesTestVectorWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestVector& item);
bool CasesTestVectorReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestVector& item);
bool CasesTestVectorWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestVector& item);

}} // namespace tl2::details

