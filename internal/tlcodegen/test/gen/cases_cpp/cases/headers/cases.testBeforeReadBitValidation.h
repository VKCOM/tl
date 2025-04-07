#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/cases.testBeforeReadBitValidation.h"

namespace tl2 { namespace details { 

void CasesTestBeforeReadBitValidationReset(::tl2::cases::TestBeforeReadBitValidation& item);

bool CasesTestBeforeReadBitValidationWriteJSON(std::ostream& s, const ::tl2::cases::TestBeforeReadBitValidation& item);
bool CasesTestBeforeReadBitValidationRead(::basictl::tl_istream & s, ::tl2::cases::TestBeforeReadBitValidation& item);
bool CasesTestBeforeReadBitValidationWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestBeforeReadBitValidation& item);
bool CasesTestBeforeReadBitValidationReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestBeforeReadBitValidation& item);
bool CasesTestBeforeReadBitValidationWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestBeforeReadBitValidation& item);

}} // namespace tl2::details

