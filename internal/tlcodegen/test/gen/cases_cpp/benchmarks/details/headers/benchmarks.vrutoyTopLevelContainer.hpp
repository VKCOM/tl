#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../types/benchmarks.vrutoyTopLevelContainer.hpp"

namespace tl2 { namespace details { 

void BenchmarksVrutoyTopLevelContainerReset(::tl2::benchmarks::VrutoyTopLevelContainer& item);
bool BenchmarksVrutoyTopLevelContainerRead(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyTopLevelContainer& item);
bool BenchmarksVrutoyTopLevelContainerWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyTopLevelContainer& item);
bool BenchmarksVrutoyTopLevelContainerReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyTopLevelContainer& item);
bool BenchmarksVrutoyTopLevelContainerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyTopLevelContainer& item);

}} // namespace tl2::details

