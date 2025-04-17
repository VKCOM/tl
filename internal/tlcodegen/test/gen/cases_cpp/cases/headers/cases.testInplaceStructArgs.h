#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases/types/cases.testInplaceStructArgs.h"

namespace tl2 { namespace details { 

void CasesTestInplaceStructArgsReset(::tl2::cases::TestInplaceStructArgs& item) noexcept;

bool CasesTestInplaceStructArgsWriteJSON(std::ostream& s, const ::tl2::cases::TestInplaceStructArgs& item) noexcept;
bool CasesTestInplaceStructArgsRead(::basictl::tl_istream & s, ::tl2::cases::TestInplaceStructArgs& item) noexcept; 
bool CasesTestInplaceStructArgsWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestInplaceStructArgs& item) noexcept;
bool CasesTestInplaceStructArgsReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestInplaceStructArgs& item);
bool CasesTestInplaceStructArgsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestInplaceStructArgs& item);

}} // namespace tl2::details

