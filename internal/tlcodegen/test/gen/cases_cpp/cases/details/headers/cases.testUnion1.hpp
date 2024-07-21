#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../types/cases.testUnion1.hpp"

namespace tl2 { namespace details { 

void CasesTestUnion1Reset(::tl2::cases::TestUnion1& item);
bool CasesTestUnion1Read(::basictl::tl_istream & s, ::tl2::cases::TestUnion1& item);
bool CasesTestUnion1Write(::basictl::tl_ostream & s, const ::tl2::cases::TestUnion1& item);
bool CasesTestUnion1ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestUnion1& item);
bool CasesTestUnion1WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestUnion1& item);

}} // namespace tl2::details

