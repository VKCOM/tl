// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "benchmarks/types/benchmarks.vrutoyTopLevelContainer.h"

namespace tlgen { namespace details { 

void BenchmarksVrutoyTopLevelContainerReset(::tlgen::benchmarks::VrutoyTopLevelContainer& item) noexcept;

bool BenchmarksVrutoyTopLevelContainerWriteJSON(std::ostream& s, const ::tlgen::benchmarks::VrutoyTopLevelContainer& item) noexcept;
bool BenchmarksVrutoyTopLevelContainerRead(::tlgen::basictl::tl_istream & s, ::tlgen::benchmarks::VrutoyTopLevelContainer& item) noexcept; 
bool BenchmarksVrutoyTopLevelContainerWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::benchmarks::VrutoyTopLevelContainer& item) noexcept;
bool BenchmarksVrutoyTopLevelContainerReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::benchmarks::VrutoyTopLevelContainer& item);
bool BenchmarksVrutoyTopLevelContainerWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::benchmarks::VrutoyTopLevelContainer& item);

}} // namespace tlgen::details

