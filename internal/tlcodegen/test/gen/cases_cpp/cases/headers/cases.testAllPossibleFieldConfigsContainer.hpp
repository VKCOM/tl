#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/cases.testAllPossibleFieldConfigsContainer.hpp"

namespace tl2 { namespace details { 

void CasesTestAllPossibleFieldConfigsContainerReset(::tl2::cases::TestAllPossibleFieldConfigsContainer& item);
bool CasesTestAllPossibleFieldConfigsContainerRead(::basictl::tl_istream & s, ::tl2::cases::TestAllPossibleFieldConfigsContainer& item);
bool CasesTestAllPossibleFieldConfigsContainerWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestAllPossibleFieldConfigsContainer& item);
bool CasesTestAllPossibleFieldConfigsContainerReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestAllPossibleFieldConfigsContainer& item);
bool CasesTestAllPossibleFieldConfigsContainerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestAllPossibleFieldConfigsContainer& item);

}} // namespace tl2::details

