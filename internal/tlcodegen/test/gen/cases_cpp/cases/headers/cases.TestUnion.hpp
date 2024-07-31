#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/cases.TestUnion.hpp"

namespace tl2 { namespace details { 

void CasesTestUnionReset(::tl2::cases::TestUnion& item);

bool CasesTestUnionWriteJSON(std::ostream & s, const ::tl2::cases::TestUnion& item);
bool CasesTestUnionReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestUnion& item);
bool CasesTestUnionWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestUnion& item);

}} // namespace tl2::details

