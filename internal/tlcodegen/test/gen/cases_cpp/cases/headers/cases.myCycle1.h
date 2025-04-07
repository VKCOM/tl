#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/cases.myCycle1.h"

namespace tl2 { namespace details { 

void CasesMyCycle1Reset(::tl2::cases::MyCycle1& item);

bool CasesMyCycle1WriteJSON(std::ostream& s, const ::tl2::cases::MyCycle1& item);
bool CasesMyCycle1Read(::basictl::tl_istream & s, ::tl2::cases::MyCycle1& item);
bool CasesMyCycle1Write(::basictl::tl_ostream & s, const ::tl2::cases::MyCycle1& item);
bool CasesMyCycle1ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::MyCycle1& item);
bool CasesMyCycle1WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::MyCycle1& item);

}} // namespace tl2::details

