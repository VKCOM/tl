#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/cases.TestEnum.h"

namespace tl2 { namespace details { 

void CasesTestEnumReset(::tl2::cases::TestEnum& item) noexcept;

bool CasesTestEnumWriteJSON(std::ostream & s, const ::tl2::cases::TestEnum& item) noexcept;
bool CasesTestEnumReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestEnum& item) noexcept;
bool CasesTestEnumWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestEnum& item) noexcept;

}} // namespace tl2::details

