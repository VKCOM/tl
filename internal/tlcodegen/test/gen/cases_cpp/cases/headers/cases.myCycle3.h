// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases/types/cases.myCycle3.h"

namespace tlgen { namespace details { 

void CasesMyCycle3Reset(::tlgen::cases::MyCycle3& item) noexcept;

bool CasesMyCycle3WriteJSON(std::ostream& s, const ::tlgen::cases::MyCycle3& item) noexcept;
bool CasesMyCycle3Read(::tlgen::basictl::tl_istream & s, ::tlgen::cases::MyCycle3& item) noexcept; 
bool CasesMyCycle3Write(::tlgen::basictl::tl_ostream & s, const ::tlgen::cases::MyCycle3& item) noexcept;
bool CasesMyCycle3ReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::cases::MyCycle3& item);
bool CasesMyCycle3WriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::cases::MyCycle3& item);

}} // namespace tlgen::details

