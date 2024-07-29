#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/benchObject.hpp"

namespace tl2 { namespace details { 

void BenchObjectReset(::tl2::BenchObject& item);
bool BenchObjectRead(::basictl::tl_istream & s, ::tl2::BenchObject& item);
bool BenchObjectWrite(::basictl::tl_ostream & s, const ::tl2::BenchObject& item);
bool BenchObjectReadBoxed(::basictl::tl_istream & s, ::tl2::BenchObject& item);
bool BenchObjectWriteBoxed(::basictl::tl_ostream & s, const ::tl2::BenchObject& item);

}} // namespace tl2::details

