#pragma once

#include "../../../basics/basictl.h"
#include "../types/benchmarks.vrutoyPositions.h"

namespace tl2 { namespace details { 

void BenchmarksVrutoyPositionsReset(::tl2::benchmarks::VrutoyPositions& item);

bool BenchmarksVrutoyPositionsWriteJSON(std::ostream& s, const ::tl2::benchmarks::VrutoyPositions& item, uint32_t nat_n);
bool BenchmarksVrutoyPositionsRead(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyPositions& item, uint32_t nat_n);
bool BenchmarksVrutoyPositionsWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyPositions& item, uint32_t nat_n);
bool BenchmarksVrutoyPositionsReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyPositions& item, uint32_t nat_n);
bool BenchmarksVrutoyPositionsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyPositions& item, uint32_t nat_n);

}} // namespace tl2::details

