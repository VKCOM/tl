#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../__common_namespace/types/vector.hpp"
#include "../types/benchmarks.vruposition.hpp"

namespace tl2 { namespace details { 

void VectorBenchmarksVruPositionReset(std::vector<::tl2::benchmarks::Vruposition>& item);
bool VectorBenchmarksVruPositionRead(::basictl::tl_istream & s, std::vector<::tl2::benchmarks::Vruposition>& item);
bool VectorBenchmarksVruPositionWrite(::basictl::tl_ostream & s, const std::vector<::tl2::benchmarks::Vruposition>& item);
bool VectorBenchmarksVruPositionReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::benchmarks::Vruposition>& item);
bool VectorBenchmarksVruPositionWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::benchmarks::Vruposition>& item);

}} // namespace tl2::details

