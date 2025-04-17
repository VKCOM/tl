#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases/types/cases.testAllPossibleFieldConfigs.h"

namespace tl2 { namespace details { 

void CasesTestAllPossibleFieldConfigsReset(::tl2::cases::TestAllPossibleFieldConfigs& item) noexcept;

bool CasesTestAllPossibleFieldConfigsWriteJSON(std::ostream& s, const ::tl2::cases::TestAllPossibleFieldConfigs& item, uint32_t nat_outer) noexcept;
bool CasesTestAllPossibleFieldConfigsRead(::basictl::tl_istream & s, ::tl2::cases::TestAllPossibleFieldConfigs& item, uint32_t nat_outer) noexcept; 
bool CasesTestAllPossibleFieldConfigsWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestAllPossibleFieldConfigs& item, uint32_t nat_outer) noexcept;
bool CasesTestAllPossibleFieldConfigsReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestAllPossibleFieldConfigs& item, uint32_t nat_outer);
bool CasesTestAllPossibleFieldConfigsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestAllPossibleFieldConfigs& item, uint32_t nat_outer);

}} // namespace tl2::details

