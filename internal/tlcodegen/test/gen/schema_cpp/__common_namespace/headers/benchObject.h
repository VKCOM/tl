#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/benchObject.h"

namespace tl2 { namespace details { 

void BenchObjectReset(::tl2::BenchObject& item) noexcept;

bool BenchObjectWriteJSON(std::ostream& s, const ::tl2::BenchObject& item) noexcept;
bool BenchObjectRead(::basictl::tl_istream & s, ::tl2::BenchObject& item) noexcept; 
bool BenchObjectWrite(::basictl::tl_ostream & s, const ::tl2::BenchObject& item) noexcept;
bool BenchObjectReadBoxed(::basictl::tl_istream & s, ::tl2::BenchObject& item);
bool BenchObjectWriteBoxed(::basictl::tl_ostream & s, const ::tl2::BenchObject& item);

}} // namespace tl2::details

