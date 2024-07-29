#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/cases.testOutFieldMask.hpp"

namespace tl2 { namespace details { 

void CasesTestOutFieldMaskReset(::tl2::cases::TestOutFieldMask& item);
bool CasesTestOutFieldMaskRead(::basictl::tl_istream & s, ::tl2::cases::TestOutFieldMask& item, uint32_t nat_f);
bool CasesTestOutFieldMaskWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestOutFieldMask& item, uint32_t nat_f);
bool CasesTestOutFieldMaskReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestOutFieldMask& item, uint32_t nat_f);
bool CasesTestOutFieldMaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestOutFieldMask& item, uint32_t nat_f);

}} // namespace tl2::details

