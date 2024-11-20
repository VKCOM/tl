#pragma once

#include "../../../basics/basictl.h"
#include "../types/benchmarks.vrutoytopLevelUnionBig.h"

namespace tl2 { namespace details { 

void BenchmarksVrutoytopLevelUnionBigReset(::tl2::benchmarks::VrutoytopLevelUnionBig& item);

bool BenchmarksVrutoytopLevelUnionBigWriteJSON(std::ostream& s, const ::tl2::benchmarks::VrutoytopLevelUnionBig& item);
bool BenchmarksVrutoytopLevelUnionBigRead(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoytopLevelUnionBig& item);
bool BenchmarksVrutoytopLevelUnionBigWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoytopLevelUnionBig& item);
bool BenchmarksVrutoytopLevelUnionBigReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoytopLevelUnionBig& item);
bool BenchmarksVrutoytopLevelUnionBigWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoytopLevelUnionBig& item);

}} // namespace tl2::details

