#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases/types/cases.testUnionContainer.h"

namespace tl2 { namespace details { 

void CasesTestUnionContainerReset(::tl2::cases::TestUnionContainer& item) noexcept;

bool CasesTestUnionContainerWriteJSON(std::ostream& s, const ::tl2::cases::TestUnionContainer& item) noexcept;
bool CasesTestUnionContainerRead(::basictl::tl_istream & s, ::tl2::cases::TestUnionContainer& item) noexcept; 
bool CasesTestUnionContainerWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestUnionContainer& item) noexcept;
bool CasesTestUnionContainerReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestUnionContainer& item);
bool CasesTestUnionContainerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestUnionContainer& item);

}} // namespace tl2::details

