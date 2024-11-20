#pragma once

#include "../../../basics/basictl.h"
#include "../types/cases.testEnumContainer.h"

namespace tl2 { namespace details { 

void CasesTestEnumContainerReset(::tl2::cases::TestEnumContainer& item);

bool CasesTestEnumContainerWriteJSON(std::ostream& s, const ::tl2::cases::TestEnumContainer& item);
bool CasesTestEnumContainerRead(::basictl::tl_istream & s, ::tl2::cases::TestEnumContainer& item);
bool CasesTestEnumContainerWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestEnumContainer& item);
bool CasesTestEnumContainerReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestEnumContainer& item);
bool CasesTestEnumContainerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestEnumContainer& item);

}} // namespace tl2::details

