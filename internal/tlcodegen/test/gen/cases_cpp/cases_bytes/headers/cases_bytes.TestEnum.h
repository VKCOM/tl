#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/cases_bytes.TestEnum.h"

namespace tl2 { namespace details { 

void CasesBytesTestEnumReset(::tl2::cases_bytes::TestEnum& item) noexcept;

bool CasesBytesTestEnumWriteJSON(std::ostream & s, const ::tl2::cases_bytes::TestEnum& item) noexcept;
bool CasesBytesTestEnumReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnum& item) noexcept;
bool CasesBytesTestEnumWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnum& item) noexcept;

}} // namespace tl2::details

