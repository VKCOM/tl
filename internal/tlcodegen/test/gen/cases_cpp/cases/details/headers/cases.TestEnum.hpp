#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../types/cases.TestEnum.hpp"

namespace tl2 { namespace details { 

void CasesTestEnumReset(::tl2::cases::TestEnum& item);
bool CasesTestEnumReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestEnum& item);
bool CasesTestEnumWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestEnum& item);

}} // namespace tl2::details

