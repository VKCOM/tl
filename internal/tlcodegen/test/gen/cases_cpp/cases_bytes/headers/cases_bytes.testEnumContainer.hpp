#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/cases_bytes.testEnumContainer.hpp"

namespace tl2 { namespace details { 

void CasesBytesTestEnumContainerReset(::tl2::cases_bytes::TestEnumContainer& item);
bool CasesBytesTestEnumContainerRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnumContainer& item);
bool CasesBytesTestEnumContainerWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnumContainer& item);
bool CasesBytesTestEnumContainerReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnumContainer& item);
bool CasesBytesTestEnumContainerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnumContainer& item);

}} // namespace tl2::details

