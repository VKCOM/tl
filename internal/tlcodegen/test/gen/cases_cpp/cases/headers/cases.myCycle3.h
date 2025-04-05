#pragma once

#include "../../basictl/io_streams.h"
#include "../types/cases.myCycle3.h"

namespace tl2 { namespace details { 

void CasesMyCycle3Reset(::tl2::cases::MyCycle3& item);

bool CasesMyCycle3WriteJSON(std::ostream& s, const ::tl2::cases::MyCycle3& item);
bool CasesMyCycle3Read(::basictl::tl_istream & s, ::tl2::cases::MyCycle3& item);
bool CasesMyCycle3Write(::basictl::tl_ostream & s, const ::tl2::cases::MyCycle3& item);
bool CasesMyCycle3ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::MyCycle3& item);
bool CasesMyCycle3WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::MyCycle3& item);

}} // namespace tl2::details

