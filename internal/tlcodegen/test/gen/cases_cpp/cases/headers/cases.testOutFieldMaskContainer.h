#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/cases.testOutFieldMaskContainer.h"

namespace tl2 { namespace details { 

void CasesTestOutFieldMaskContainerReset(::tl2::cases::TestOutFieldMaskContainer& item);

bool CasesTestOutFieldMaskContainerWriteJSON(std::ostream& s, const ::tl2::cases::TestOutFieldMaskContainer& item);
bool CasesTestOutFieldMaskContainerRead(::basictl::tl_istream & s, ::tl2::cases::TestOutFieldMaskContainer& item);
bool CasesTestOutFieldMaskContainerWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestOutFieldMaskContainer& item);
bool CasesTestOutFieldMaskContainerReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestOutFieldMaskContainer& item);
bool CasesTestOutFieldMaskContainerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestOutFieldMaskContainer& item);

}} // namespace tl2::details

