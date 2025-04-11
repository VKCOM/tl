#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/cases.testUnion2.h"

namespace tl2 { namespace details { 

void CasesTestUnion2Reset(::tl2::cases::TestUnion2& item) noexcept;

bool CasesTestUnion2WriteJSON(std::ostream& s, const ::tl2::cases::TestUnion2& item) noexcept;
bool CasesTestUnion2Read(::basictl::tl_istream & s, ::tl2::cases::TestUnion2& item) noexcept; 
bool CasesTestUnion2Write(::basictl::tl_ostream & s, const ::tl2::cases::TestUnion2& item) noexcept;
bool CasesTestUnion2ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestUnion2& item);
bool CasesTestUnion2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestUnion2& item);

}} // namespace tl2::details

