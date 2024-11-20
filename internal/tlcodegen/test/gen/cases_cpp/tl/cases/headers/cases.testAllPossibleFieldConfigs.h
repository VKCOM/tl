#pragma once

#include "../../../basics/basictl.h"
#include "../types/cases.testAllPossibleFieldConfigs.h"

namespace tl2 { namespace details { 

void CasesTestAllPossibleFieldConfigsReset(::tl2::cases::TestAllPossibleFieldConfigs& item);

bool CasesTestAllPossibleFieldConfigsWriteJSON(std::ostream& s, const ::tl2::cases::TestAllPossibleFieldConfigs& item, uint32_t nat_outer);
bool CasesTestAllPossibleFieldConfigsRead(::basictl::tl_istream & s, ::tl2::cases::TestAllPossibleFieldConfigs& item, uint32_t nat_outer);
bool CasesTestAllPossibleFieldConfigsWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestAllPossibleFieldConfigs& item, uint32_t nat_outer);
bool CasesTestAllPossibleFieldConfigsReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestAllPossibleFieldConfigs& item, uint32_t nat_outer);
bool CasesTestAllPossibleFieldConfigsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestAllPossibleFieldConfigs& item, uint32_t nat_outer);

}} // namespace tl2::details

