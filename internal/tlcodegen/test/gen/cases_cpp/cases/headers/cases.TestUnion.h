#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/cases.TestUnion.h"

namespace tl2 { namespace details { 

void CasesTestUnionReset(::tl2::cases::TestUnion& item) noexcept;

bool CasesTestUnionWriteJSON(std::ostream & s, const ::tl2::cases::TestUnion& item) noexcept;
bool CasesTestUnionReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestUnion& item) noexcept;
bool CasesTestUnionWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestUnion& item) noexcept;

}} // namespace tl2::details

