#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/cases.testMaybe.h"

namespace tl2 { namespace details { 

void CasesTestMaybeReset(::tl2::cases::TestMaybe& item) noexcept;

bool CasesTestMaybeWriteJSON(std::ostream& s, const ::tl2::cases::TestMaybe& item) noexcept;
bool CasesTestMaybeRead(::basictl::tl_istream & s, ::tl2::cases::TestMaybe& item) noexcept; 
bool CasesTestMaybeWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestMaybe& item) noexcept;
bool CasesTestMaybeReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestMaybe& item);
bool CasesTestMaybeWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestMaybe& item);

}} // namespace tl2::details

