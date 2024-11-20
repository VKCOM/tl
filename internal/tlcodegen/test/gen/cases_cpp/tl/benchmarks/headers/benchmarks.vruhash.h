#pragma once

#include "../../../basics/basictl.h"
#include "../types/benchmarks.vruhash.h"

namespace tl2 { namespace details { 

void BenchmarksVruHashReset(::tl2::benchmarks::Vruhash& item);

bool BenchmarksVruHashWriteJSON(std::ostream& s, const ::tl2::benchmarks::Vruhash& item);
bool BenchmarksVruHashRead(::basictl::tl_istream & s, ::tl2::benchmarks::Vruhash& item);
bool BenchmarksVruHashWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::Vruhash& item);
bool BenchmarksVruHashReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::Vruhash& item);
bool BenchmarksVruHashWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::Vruhash& item);

}} // namespace tl2::details

