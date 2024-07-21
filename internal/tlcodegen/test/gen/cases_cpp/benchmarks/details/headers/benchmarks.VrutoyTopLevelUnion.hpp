#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../types/benchmarks.VrutoyTopLevelUnion.hpp"

namespace tl2 { namespace details { 

void BenchmarksVrutoyTopLevelUnionReset(::tl2::benchmarks::VrutoyTopLevelUnion& item);
bool BenchmarksVrutoyTopLevelUnionReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyTopLevelUnion& item);
bool BenchmarksVrutoyTopLevelUnionWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyTopLevelUnion& item);

}} // namespace tl2::details

