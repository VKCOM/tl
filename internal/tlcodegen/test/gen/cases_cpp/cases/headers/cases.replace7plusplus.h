#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases/types/cases.replace7plusplus.h"

namespace tl2 { namespace details { 

void CasesReplace7plusplusReset(::tl2::cases::Replace7plusplus& item) noexcept;

bool CasesReplace7plusplusWriteJSON(std::ostream& s, const ::tl2::cases::Replace7plusplus& item) noexcept;
bool CasesReplace7plusplusRead(::basictl::tl_istream & s, ::tl2::cases::Replace7plusplus& item) noexcept; 
bool CasesReplace7plusplusWrite(::basictl::tl_ostream & s, const ::tl2::cases::Replace7plusplus& item) noexcept;
bool CasesReplace7plusplusReadBoxed(::basictl::tl_istream & s, ::tl2::cases::Replace7plusplus& item);
bool CasesReplace7plusplusWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::Replace7plusplus& item);

}} // namespace tl2::details

