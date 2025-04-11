#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/cases.myCycle2.h"

namespace tl2 { namespace details { 

void CasesMyCycle2Reset(::tl2::cases::MyCycle2& item) noexcept;

bool CasesMyCycle2WriteJSON(std::ostream& s, const ::tl2::cases::MyCycle2& item) noexcept;
bool CasesMyCycle2Read(::basictl::tl_istream & s, ::tl2::cases::MyCycle2& item) noexcept; 
bool CasesMyCycle2Write(::basictl::tl_ostream & s, const ::tl2::cases::MyCycle2& item) noexcept;
bool CasesMyCycle2ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::MyCycle2& item);
bool CasesMyCycle2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::MyCycle2& item);

}} // namespace tl2::details

