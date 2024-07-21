#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../types/cases.testArray.hpp"

namespace tl2 { namespace details { 

void CasesTestArrayReset(::tl2::cases::TestArray& item);
bool CasesTestArrayRead(::basictl::tl_istream & s, ::tl2::cases::TestArray& item);
bool CasesTestArrayWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestArray& item);
bool CasesTestArrayReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestArray& item);
bool CasesTestArrayWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestArray& item);

}} // namespace tl2::details

