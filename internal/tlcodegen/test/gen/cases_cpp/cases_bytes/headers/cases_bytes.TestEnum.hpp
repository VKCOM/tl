#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/cases_bytes.TestEnum.hpp"

namespace tl2 { namespace details { 

void CasesBytesTestEnumReset(::tl2::cases_bytes::TestEnum& item);

bool CasesBytesTestEnumWriteJSON(std::ostream & s, const ::tl2::cases_bytes::TestEnum& item);
bool CasesBytesTestEnumReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnum& item);
bool CasesBytesTestEnumWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnum& item);

}} // namespace tl2::details

