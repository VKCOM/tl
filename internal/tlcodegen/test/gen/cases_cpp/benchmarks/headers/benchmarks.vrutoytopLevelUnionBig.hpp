#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/benchmarks.vrutoytopLevelUnionBig.hpp"

namespace tl2 { namespace details { 

void BenchmarksVrutoytopLevelUnionBigReset(::tl2::benchmarks::VrutoytopLevelUnionBig& item);

bool BenchmarksVrutoytopLevelUnionBigWriteJSON(std::ostream& s, const ::tl2::benchmarks::VrutoytopLevelUnionBig& item);
bool BenchmarksVrutoytopLevelUnionBigRead(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoytopLevelUnionBig& item);
bool BenchmarksVrutoytopLevelUnionBigWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoytopLevelUnionBig& item);
bool BenchmarksVrutoytopLevelUnionBigReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoytopLevelUnionBig& item);
bool BenchmarksVrutoytopLevelUnionBigWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoytopLevelUnionBig& item);

}} // namespace tl2::details

