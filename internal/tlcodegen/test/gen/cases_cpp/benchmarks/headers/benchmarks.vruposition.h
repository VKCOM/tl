#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/benchmarks.vruposition.h"

namespace tl2 { namespace details { 

void BenchmarksVruPositionReset(::tl2::benchmarks::Vruposition& item) noexcept;

bool BenchmarksVruPositionWriteJSON(std::ostream& s, const ::tl2::benchmarks::Vruposition& item) noexcept;
bool BenchmarksVruPositionRead(::basictl::tl_istream & s, ::tl2::benchmarks::Vruposition& item) noexcept; 
bool BenchmarksVruPositionWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::Vruposition& item) noexcept;
bool BenchmarksVruPositionReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::Vruposition& item);
bool BenchmarksVruPositionWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::Vruposition& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinTupleBenchmarksVruPositionReset(std::vector<::tl2::benchmarks::Vruposition>& item);

bool BuiltinTupleBenchmarksVruPositionWriteJSON(std::ostream & s, const std::vector<::tl2::benchmarks::Vruposition>& item, uint32_t nat_n);
bool BuiltinTupleBenchmarksVruPositionRead(::basictl::tl_istream & s, std::vector<::tl2::benchmarks::Vruposition>& item, uint32_t nat_n);
bool BuiltinTupleBenchmarksVruPositionWrite(::basictl::tl_ostream & s, const std::vector<::tl2::benchmarks::Vruposition>& item, uint32_t nat_n);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorBenchmarksVruPositionReset(std::vector<::tl2::benchmarks::Vruposition>& item);

bool BuiltinVectorBenchmarksVruPositionWriteJSON(std::ostream & s, const std::vector<::tl2::benchmarks::Vruposition>& item);
bool BuiltinVectorBenchmarksVruPositionRead(::basictl::tl_istream & s, std::vector<::tl2::benchmarks::Vruposition>& item);
bool BuiltinVectorBenchmarksVruPositionWrite(::basictl::tl_ostream & s, const std::vector<::tl2::benchmarks::Vruposition>& item);

}} // namespace tl2::details

