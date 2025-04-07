#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/cases.testInplaceStructArgs2.h"

namespace tl2 { namespace details { 

void CasesTestInplaceStructArgs2Reset(::tl2::cases::TestInplaceStructArgs2& item);

bool CasesTestInplaceStructArgs2WriteJSON(std::ostream& s, const ::tl2::cases::TestInplaceStructArgs2& item);
bool CasesTestInplaceStructArgs2Read(::basictl::tl_istream & s, ::tl2::cases::TestInplaceStructArgs2& item);
bool CasesTestInplaceStructArgs2Write(::basictl::tl_ostream & s, const ::tl2::cases::TestInplaceStructArgs2& item);
bool CasesTestInplaceStructArgs2ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestInplaceStructArgs2& item);
bool CasesTestInplaceStructArgs2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestInplaceStructArgs2& item);

}} // namespace tl2::details

