#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/cases.testRecursiveFieldMask.hpp"

namespace tl2 { namespace details { 

void CasesTestRecursiveFieldmaskReset(::tl2::cases::TestRecursiveFieldMask& item);

bool CasesTestRecursiveFieldmaskWriteJSON(std::ostream& s, const ::tl2::cases::TestRecursiveFieldMask& item);
bool CasesTestRecursiveFieldmaskRead(::basictl::tl_istream & s, ::tl2::cases::TestRecursiveFieldMask& item);
bool CasesTestRecursiveFieldmaskWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestRecursiveFieldMask& item);
bool CasesTestRecursiveFieldmaskReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestRecursiveFieldMask& item);
bool CasesTestRecursiveFieldmaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestRecursiveFieldMask& item);

}} // namespace tl2::details

