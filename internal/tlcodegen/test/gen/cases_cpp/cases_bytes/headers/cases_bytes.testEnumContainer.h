// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases_bytes/types/cases_bytes.testEnumContainer.h"

namespace tl2 { namespace details { 

void CasesBytesTestEnumContainerReset(::tl2::cases_bytes::TestEnumContainer& item) noexcept;

bool CasesBytesTestEnumContainerWriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestEnumContainer& item) noexcept;
bool CasesBytesTestEnumContainerRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnumContainer& item) noexcept; 
bool CasesBytesTestEnumContainerWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnumContainer& item) noexcept;
bool CasesBytesTestEnumContainerReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnumContainer& item);
bool CasesBytesTestEnumContainerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnumContainer& item);

}} // namespace tl2::details

