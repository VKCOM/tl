#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/cases.testAllPossibleFieldConfigs.hpp"

namespace tl2 { namespace details { 

void CasesTestAllPossibleFieldConfigsReset(::tl2::cases::TestAllPossibleFieldConfigs& item);
bool CasesTestAllPossibleFieldConfigsRead(::basictl::tl_istream & s, ::tl2::cases::TestAllPossibleFieldConfigs& item, uint32_t nat_outer);
bool CasesTestAllPossibleFieldConfigsWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestAllPossibleFieldConfigs& item, uint32_t nat_outer);
bool CasesTestAllPossibleFieldConfigsReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestAllPossibleFieldConfigs& item, uint32_t nat_outer);
bool CasesTestAllPossibleFieldConfigsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestAllPossibleFieldConfigs& item, uint32_t nat_outer);

}} // namespace tl2::details

