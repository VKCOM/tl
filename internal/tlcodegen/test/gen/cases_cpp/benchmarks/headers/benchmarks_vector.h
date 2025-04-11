#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../../__common_namespace/types/vector.h"
#include "../types/benchmarks.vruposition.h"

namespace tl2 { namespace details { 

void VectorBenchmarksVruPositionReset(std::vector<::tl2::benchmarks::Vruposition>& item) noexcept;

bool VectorBenchmarksVruPositionWriteJSON(std::ostream& s, const std::vector<::tl2::benchmarks::Vruposition>& item) noexcept;
bool VectorBenchmarksVruPositionRead(::basictl::tl_istream & s, std::vector<::tl2::benchmarks::Vruposition>& item) noexcept; 
bool VectorBenchmarksVruPositionWrite(::basictl::tl_ostream & s, const std::vector<::tl2::benchmarks::Vruposition>& item) noexcept;
bool VectorBenchmarksVruPositionReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::benchmarks::Vruposition>& item);
bool VectorBenchmarksVruPositionWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::benchmarks::Vruposition>& item);

}} // namespace tl2::details

