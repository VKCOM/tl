#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/cases.testMaybe.hpp"

namespace tl2 { namespace details { 

void CasesTestMaybeReset(::tl2::cases::TestMaybe& item);

bool CasesTestMaybeWriteJSON(std::ostream& s, const ::tl2::cases::TestMaybe& item);
bool CasesTestMaybeRead(::basictl::tl_istream & s, ::tl2::cases::TestMaybe& item);
bool CasesTestMaybeWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestMaybe& item);
bool CasesTestMaybeReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestMaybe& item);
bool CasesTestMaybeWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestMaybe& item);

}} // namespace tl2::details

