// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases/types/cases.testUnionContainer.h"

namespace tlgen { namespace details { 

void CasesTestUnionContainerReset(::tlgen::cases::TestUnionContainer& item) noexcept;

bool CasesTestUnionContainerWriteJSON(std::ostream& s, const ::tlgen::cases::TestUnionContainer& item) noexcept;
bool CasesTestUnionContainerRead(::tlgen::basictl::tl_istream & s, ::tlgen::cases::TestUnionContainer& item) noexcept; 
bool CasesTestUnionContainerWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::cases::TestUnionContainer& item) noexcept;
bool CasesTestUnionContainerReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::cases::TestUnionContainer& item);
bool CasesTestUnionContainerWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::cases::TestUnionContainer& item);

}} // namespace tlgen::details

