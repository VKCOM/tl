#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases/types/cases.testLocalFieldmask.h"

namespace tl2 { namespace details { 

void CasesTestLocalFieldmaskReset(::tl2::cases::TestLocalFieldmask& item) noexcept;

bool CasesTestLocalFieldmaskWriteJSON(std::ostream& s, const ::tl2::cases::TestLocalFieldmask& item) noexcept;
bool CasesTestLocalFieldmaskRead(::basictl::tl_istream & s, ::tl2::cases::TestLocalFieldmask& item) noexcept; 
bool CasesTestLocalFieldmaskWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestLocalFieldmask& item) noexcept;
bool CasesTestLocalFieldmaskReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestLocalFieldmask& item);
bool CasesTestLocalFieldmaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestLocalFieldmask& item);

}} // namespace tl2::details

