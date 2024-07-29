#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/benchmarks.vruposition.hpp"

namespace tl2 { namespace details { 

void BenchmarksVruPositionReset(::tl2::benchmarks::Vruposition& item);
bool BenchmarksVruPositionRead(::basictl::tl_istream & s, ::tl2::benchmarks::Vruposition& item);
bool BenchmarksVruPositionWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::Vruposition& item);
bool BenchmarksVruPositionReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::Vruposition& item);
bool BenchmarksVruPositionWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::Vruposition& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinTupleBenchmarksVruPositionReset(std::vector<::tl2::benchmarks::Vruposition>& item);
bool BuiltinTupleBenchmarksVruPositionRead(::basictl::tl_istream & s, std::vector<::tl2::benchmarks::Vruposition>& item, uint32_t nat_n);
bool BuiltinTupleBenchmarksVruPositionWrite(::basictl::tl_ostream & s, const std::vector<::tl2::benchmarks::Vruposition>& item, uint32_t nat_n);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorBenchmarksVruPositionReset(std::vector<::tl2::benchmarks::Vruposition>& item);
bool BuiltinVectorBenchmarksVruPositionRead(::basictl::tl_istream & s, std::vector<::tl2::benchmarks::Vruposition>& item);
bool BuiltinVectorBenchmarksVruPositionWrite(::basictl::tl_ostream & s, const std::vector<::tl2::benchmarks::Vruposition>& item);

}} // namespace tl2::details

