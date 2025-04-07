#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/benchmarks.VrutoyTopLevelUnion.h"

namespace tl2 { namespace details { 

void BenchmarksVrutoyTopLevelUnionReset(::tl2::benchmarks::VrutoyTopLevelUnion& item);

bool BenchmarksVrutoyTopLevelUnionWriteJSON(std::ostream & s, const ::tl2::benchmarks::VrutoyTopLevelUnion& item);
bool BenchmarksVrutoyTopLevelUnionReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyTopLevelUnion& item);
bool BenchmarksVrutoyTopLevelUnionWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyTopLevelUnion& item);

}} // namespace tl2::details

