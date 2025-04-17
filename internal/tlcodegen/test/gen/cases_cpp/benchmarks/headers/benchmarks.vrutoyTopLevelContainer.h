#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "benchmarks/types/benchmarks.vrutoyTopLevelContainer.h"

namespace tl2 { namespace details { 

void BenchmarksVrutoyTopLevelContainerReset(::tl2::benchmarks::VrutoyTopLevelContainer& item) noexcept;

bool BenchmarksVrutoyTopLevelContainerWriteJSON(std::ostream& s, const ::tl2::benchmarks::VrutoyTopLevelContainer& item) noexcept;
bool BenchmarksVrutoyTopLevelContainerRead(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyTopLevelContainer& item) noexcept; 
bool BenchmarksVrutoyTopLevelContainerWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyTopLevelContainer& item) noexcept;
bool BenchmarksVrutoyTopLevelContainerReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyTopLevelContainer& item);
bool BenchmarksVrutoyTopLevelContainerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyTopLevelContainer& item);

}} // namespace tl2::details

