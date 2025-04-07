#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/cases.testAllPossibleFieldConfigsContainer.h"

namespace tl2 { namespace details { 

void CasesTestAllPossibleFieldConfigsContainerReset(::tl2::cases::TestAllPossibleFieldConfigsContainer& item);

bool CasesTestAllPossibleFieldConfigsContainerWriteJSON(std::ostream& s, const ::tl2::cases::TestAllPossibleFieldConfigsContainer& item);
bool CasesTestAllPossibleFieldConfigsContainerRead(::basictl::tl_istream & s, ::tl2::cases::TestAllPossibleFieldConfigsContainer& item);
bool CasesTestAllPossibleFieldConfigsContainerWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestAllPossibleFieldConfigsContainer& item);
bool CasesTestAllPossibleFieldConfigsContainerReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestAllPossibleFieldConfigsContainer& item);
bool CasesTestAllPossibleFieldConfigsContainerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestAllPossibleFieldConfigsContainer& item);

}} // namespace tl2::details

