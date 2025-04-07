#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/cases.testOutFieldMask.h"

namespace tl2 { namespace details { 

void CasesTestOutFieldMaskReset(::tl2::cases::TestOutFieldMask& item);

bool CasesTestOutFieldMaskWriteJSON(std::ostream& s, const ::tl2::cases::TestOutFieldMask& item, uint32_t nat_f);
bool CasesTestOutFieldMaskRead(::basictl::tl_istream & s, ::tl2::cases::TestOutFieldMask& item, uint32_t nat_f);
bool CasesTestOutFieldMaskWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestOutFieldMask& item, uint32_t nat_f);
bool CasesTestOutFieldMaskReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestOutFieldMask& item, uint32_t nat_f);
bool CasesTestOutFieldMaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestOutFieldMask& item, uint32_t nat_f);

}} // namespace tl2::details

