#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/cases.testUnionContainer.hpp"

namespace tl2 { namespace details { 

void CasesTestUnionContainerReset(::tl2::cases::TestUnionContainer& item);

bool CasesTestUnionContainerWriteJSON(std::ostream& s, const ::tl2::cases::TestUnionContainer& item);
bool CasesTestUnionContainerRead(::basictl::tl_istream & s, ::tl2::cases::TestUnionContainer& item);
bool CasesTestUnionContainerWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestUnionContainer& item);
bool CasesTestUnionContainerReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestUnionContainer& item);
bool CasesTestUnionContainerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestUnionContainer& item);

}} // namespace tl2::details

