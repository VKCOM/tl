#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../../__common_namespace/types/boolStat.h"

namespace tl2 { namespace details { 

void BoolStatReset(::tl2::BoolStat& item) noexcept;

bool BoolStatWriteJSON(std::ostream& s, const ::tl2::BoolStat& item) noexcept;
bool BoolStatRead(::basictl::tl_istream & s, ::tl2::BoolStat& item) noexcept; 
bool BoolStatWrite(::basictl::tl_ostream & s, const ::tl2::BoolStat& item) noexcept;
bool BoolStatReadBoxed(::basictl::tl_istream & s, ::tl2::BoolStat& item);
bool BoolStatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoolStat& item);

}} // namespace tl2::details

