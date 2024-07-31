#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/cases.testUnion2.hpp"

namespace tl2 { namespace details { 

void CasesTestUnion2Reset(::tl2::cases::TestUnion2& item);

bool CasesTestUnion2WriteJSON(std::ostream& s, const ::tl2::cases::TestUnion2& item);
bool CasesTestUnion2Read(::basictl::tl_istream & s, ::tl2::cases::TestUnion2& item);
bool CasesTestUnion2Write(::basictl::tl_ostream & s, const ::tl2::cases::TestUnion2& item);
bool CasesTestUnion2ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestUnion2& item);
bool CasesTestUnion2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestUnion2& item);

}} // namespace tl2::details

