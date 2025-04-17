#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "benchmarks/types/benchmarks.vrutoytopLevelUnionBig.h"

namespace tl2 { namespace details { 

void BenchmarksVrutoytopLevelUnionBigReset(::tl2::benchmarks::VrutoytopLevelUnionBig& item) noexcept;

bool BenchmarksVrutoytopLevelUnionBigWriteJSON(std::ostream& s, const ::tl2::benchmarks::VrutoytopLevelUnionBig& item) noexcept;
bool BenchmarksVrutoytopLevelUnionBigRead(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoytopLevelUnionBig& item) noexcept; 
bool BenchmarksVrutoytopLevelUnionBigWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoytopLevelUnionBig& item) noexcept;
bool BenchmarksVrutoytopLevelUnionBigReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoytopLevelUnionBig& item);
bool BenchmarksVrutoytopLevelUnionBigWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoytopLevelUnionBig& item);

}} // namespace tl2::details

