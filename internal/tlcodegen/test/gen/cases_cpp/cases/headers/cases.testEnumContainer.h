#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases/types/cases.testEnumContainer.h"

namespace tl2 { namespace details { 

void CasesTestEnumContainerReset(::tl2::cases::TestEnumContainer& item) noexcept;

bool CasesTestEnumContainerWriteJSON(std::ostream& s, const ::tl2::cases::TestEnumContainer& item) noexcept;
bool CasesTestEnumContainerRead(::basictl::tl_istream & s, ::tl2::cases::TestEnumContainer& item) noexcept; 
bool CasesTestEnumContainerWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestEnumContainer& item) noexcept;
bool CasesTestEnumContainerReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestEnumContainer& item);
bool CasesTestEnumContainerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestEnumContainer& item);

}} // namespace tl2::details

