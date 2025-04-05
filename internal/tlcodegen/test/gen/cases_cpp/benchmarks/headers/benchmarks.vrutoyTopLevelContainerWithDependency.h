#pragma once

#include "../../basictl/io_streams.h"
#include "../types/benchmarks.vrutoyTopLevelContainerWithDependency.h"

namespace tl2 { namespace details { 

void BenchmarksVrutoyTopLevelContainerWithDependencyReset(::tl2::benchmarks::VrutoyTopLevelContainerWithDependency& item);

bool BenchmarksVrutoyTopLevelContainerWithDependencyWriteJSON(std::ostream& s, const ::tl2::benchmarks::VrutoyTopLevelContainerWithDependency& item);
bool BenchmarksVrutoyTopLevelContainerWithDependencyRead(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyTopLevelContainerWithDependency& item);
bool BenchmarksVrutoyTopLevelContainerWithDependencyWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyTopLevelContainerWithDependency& item);
bool BenchmarksVrutoyTopLevelContainerWithDependencyReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyTopLevelContainerWithDependency& item);
bool BenchmarksVrutoyTopLevelContainerWithDependencyWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyTopLevelContainerWithDependency& item);

}} // namespace tl2::details

