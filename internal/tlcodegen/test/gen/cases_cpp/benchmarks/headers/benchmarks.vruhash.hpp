#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/benchmarks.vruhash.hpp"

namespace tl2 { namespace details { 

void BenchmarksVruHashReset(::tl2::benchmarks::Vruhash& item);

bool BenchmarksVruHashWriteJSON(std::ostream& s, const ::tl2::benchmarks::Vruhash& item);
bool BenchmarksVruHashRead(::basictl::tl_istream & s, ::tl2::benchmarks::Vruhash& item);
bool BenchmarksVruHashWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::Vruhash& item);
bool BenchmarksVruHashReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::Vruhash& item);
bool BenchmarksVruHashWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::Vruhash& item);

}} // namespace tl2::details

