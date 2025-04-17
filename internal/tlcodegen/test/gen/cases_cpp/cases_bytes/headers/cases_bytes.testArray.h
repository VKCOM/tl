#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases_bytes/types/cases_bytes.testArray.h"

namespace tl2 { namespace details { 

void CasesBytesTestArrayReset(::tl2::cases_bytes::TestArray& item) noexcept;

bool CasesBytesTestArrayWriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestArray& item) noexcept;
bool CasesBytesTestArrayRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestArray& item) noexcept; 
bool CasesBytesTestArrayWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestArray& item) noexcept;
bool CasesBytesTestArrayReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestArray& item);
bool CasesBytesTestArrayWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestArray& item);

}} // namespace tl2::details

