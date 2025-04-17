#pragma once

#include "../../basictl/io_streams.h"
#include "../types/cases_bytes.testEnumContainer.h"

namespace tl2 { namespace details { 

void CasesBytesTestEnumContainerReset(::tl2::cases_bytes::TestEnumContainer& item);

bool CasesBytesTestEnumContainerWriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestEnumContainer& item);
bool CasesBytesTestEnumContainerRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnumContainer& item);
bool CasesBytesTestEnumContainerWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnumContainer& item);
bool CasesBytesTestEnumContainerReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnumContainer& item);
bool CasesBytesTestEnumContainerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnumContainer& item);

}} // namespace tl2::details

