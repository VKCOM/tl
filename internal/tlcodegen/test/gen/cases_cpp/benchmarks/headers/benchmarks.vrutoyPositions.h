#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "benchmarks/types/benchmarks.vrutoyPositions.h"

namespace tl2 { namespace details { 

void BenchmarksVrutoyPositionsReset(::tl2::benchmarks::VrutoyPositions& item) noexcept;

bool BenchmarksVrutoyPositionsWriteJSON(std::ostream& s, const ::tl2::benchmarks::VrutoyPositions& item, uint32_t nat_n) noexcept;
bool BenchmarksVrutoyPositionsRead(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyPositions& item, uint32_t nat_n) noexcept; 
bool BenchmarksVrutoyPositionsWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyPositions& item, uint32_t nat_n) noexcept;
bool BenchmarksVrutoyPositionsReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyPositions& item, uint32_t nat_n);
bool BenchmarksVrutoyPositionsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyPositions& item, uint32_t nat_n);

}} // namespace tl2::details

