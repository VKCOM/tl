#pragma once

#include "../../../basics/basictl.h"
#include "../types/cases.TestEnum.h"

namespace tl2 { namespace details { 

void CasesTestEnumReset(::tl2::cases::TestEnum& item);

bool CasesTestEnumWriteJSON(std::ostream & s, const ::tl2::cases::TestEnum& item);
bool CasesTestEnumReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestEnum& item);
bool CasesTestEnumWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestEnum& item);

}} // namespace tl2::details

