// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "benchmarks/types/benchmarks.vruposition.h"

namespace tlgen { namespace details { 

void BenchmarksVruPositionReset(::tlgen::benchmarks::Vruposition& item) noexcept;

bool BenchmarksVruPositionWriteJSON(std::ostream& s, const ::tlgen::benchmarks::Vruposition& item) noexcept;
bool BenchmarksVruPositionRead(::tlgen::basictl::tl_istream & s, ::tlgen::benchmarks::Vruposition& item) noexcept; 
bool BenchmarksVruPositionWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::benchmarks::Vruposition& item) noexcept;
bool BenchmarksVruPositionReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::benchmarks::Vruposition& item);
bool BenchmarksVruPositionWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::benchmarks::Vruposition& item);

}} // namespace tlgen::details

namespace tlgen { namespace details { 

void BuiltinTupleBenchmarksVruPositionReset(std::vector<::tlgen::benchmarks::Vruposition>& item);

bool BuiltinTupleBenchmarksVruPositionWriteJSON(std::ostream & s, const std::vector<::tlgen::benchmarks::Vruposition>& item, [[maybe_unused]] uint32_t nat_n);
bool BuiltinTupleBenchmarksVruPositionRead(::tlgen::basictl::tl_istream & s, std::vector<::tlgen::benchmarks::Vruposition>& item, [[maybe_unused]] uint32_t nat_n);
bool BuiltinTupleBenchmarksVruPositionWrite(::tlgen::basictl::tl_ostream & s, const std::vector<::tlgen::benchmarks::Vruposition>& item, [[maybe_unused]] uint32_t nat_n);

}} // namespace tlgen::details

namespace tlgen { namespace details { 

void BuiltinVectorBenchmarksVruPositionReset(std::vector<::tlgen::benchmarks::Vruposition>& item);

bool BuiltinVectorBenchmarksVruPositionWriteJSON(std::ostream & s, const std::vector<::tlgen::benchmarks::Vruposition>& item);
bool BuiltinVectorBenchmarksVruPositionRead(::tlgen::basictl::tl_istream & s, std::vector<::tlgen::benchmarks::Vruposition>& item);
bool BuiltinVectorBenchmarksVruPositionWrite(::tlgen::basictl::tl_ostream & s, const std::vector<::tlgen::benchmarks::Vruposition>& item);

}} // namespace tlgen::details

