// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases_bytes/types/cases_bytes.TestEnum.h"

namespace tlgen { namespace details { 

void CasesBytesTestEnumReset(::tlgen::cases_bytes::TestEnum& item) noexcept;

bool CasesBytesTestEnumWriteJSON(std::ostream & s, const ::tlgen::cases_bytes::TestEnum& item) noexcept;
bool CasesBytesTestEnumReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::cases_bytes::TestEnum& item) noexcept;
bool CasesBytesTestEnumWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::cases_bytes::TestEnum& item) noexcept;

}} // namespace tlgen::details

