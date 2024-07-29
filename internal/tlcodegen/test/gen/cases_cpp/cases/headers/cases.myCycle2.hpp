#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/cases.myCycle2.hpp"

namespace tl2 { namespace details { 

void CasesMyCycle2Reset(::tl2::cases::MyCycle2& item);
bool CasesMyCycle2Read(::basictl::tl_istream & s, ::tl2::cases::MyCycle2& item);
bool CasesMyCycle2Write(::basictl::tl_ostream & s, const ::tl2::cases::MyCycle2& item);
bool CasesMyCycle2ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::MyCycle2& item);
bool CasesMyCycle2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::MyCycle2& item);

}} // namespace tl2::details

