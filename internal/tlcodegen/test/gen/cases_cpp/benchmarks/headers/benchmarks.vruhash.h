#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/benchmarks.vruhash.h"

namespace tl2 { namespace details { 

void BenchmarksVruHashReset(::tl2::benchmarks::Vruhash& item) noexcept;

bool BenchmarksVruHashWriteJSON(std::ostream& s, const ::tl2::benchmarks::Vruhash& item) noexcept;
bool BenchmarksVruHashRead(::basictl::tl_istream & s, ::tl2::benchmarks::Vruhash& item) noexcept; 
bool BenchmarksVruHashWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::Vruhash& item) noexcept;
bool BenchmarksVruHashReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::Vruhash& item);
bool BenchmarksVruHashWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::Vruhash& item);

}} // namespace tl2::details

