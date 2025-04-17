#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases/types/cases.testArray.h"

namespace tl2 { namespace details { 

void CasesTestArrayReset(::tl2::cases::TestArray& item) noexcept;

bool CasesTestArrayWriteJSON(std::ostream& s, const ::tl2::cases::TestArray& item) noexcept;
bool CasesTestArrayRead(::basictl::tl_istream & s, ::tl2::cases::TestArray& item) noexcept; 
bool CasesTestArrayWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestArray& item) noexcept;
bool CasesTestArrayReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestArray& item);
bool CasesTestArrayWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestArray& item);

}} // namespace tl2::details

