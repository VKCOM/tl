#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/cases.testRecursiveFieldMask.h"

namespace tl2 { namespace details { 

void CasesTestRecursiveFieldmaskReset(::tl2::cases::TestRecursiveFieldMask& item);

bool CasesTestRecursiveFieldmaskWriteJSON(std::ostream& s, const ::tl2::cases::TestRecursiveFieldMask& item);
bool CasesTestRecursiveFieldmaskRead(::basictl::tl_istream & s, ::tl2::cases::TestRecursiveFieldMask& item);
bool CasesTestRecursiveFieldmaskWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestRecursiveFieldMask& item);
bool CasesTestRecursiveFieldmaskReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestRecursiveFieldMask& item);
bool CasesTestRecursiveFieldmaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestRecursiveFieldMask& item);

}} // namespace tl2::details

