#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/benchmarks.vrutoyPositions.hpp"

namespace tl2 { namespace details { 

void BenchmarksVrutoyPositionsReset(::tl2::benchmarks::VrutoyPositions& item);
bool BenchmarksVrutoyPositionsRead(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyPositions& item, uint32_t nat_n);
bool BenchmarksVrutoyPositionsWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyPositions& item, uint32_t nat_n);
bool BenchmarksVrutoyPositionsReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyPositions& item, uint32_t nat_n);
bool BenchmarksVrutoyPositionsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyPositions& item, uint32_t nat_n);

}} // namespace tl2::details

