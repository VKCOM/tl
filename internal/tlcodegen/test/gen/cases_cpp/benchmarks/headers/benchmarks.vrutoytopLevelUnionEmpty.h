#pragma once

#include "../../basictl/io_streams.h"
#include "../types/benchmarks.vrutoytopLevelUnionEmpty.h"

namespace tl2 { namespace details { 

void BenchmarksVrutoytopLevelUnionEmptyReset(::tl2::benchmarks::VrutoytopLevelUnionEmpty& item);

bool BenchmarksVrutoytopLevelUnionEmptyWriteJSON(std::ostream& s, const ::tl2::benchmarks::VrutoytopLevelUnionEmpty& item);
bool BenchmarksVrutoytopLevelUnionEmptyRead(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoytopLevelUnionEmpty& item);
bool BenchmarksVrutoytopLevelUnionEmptyWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoytopLevelUnionEmpty& item);
bool BenchmarksVrutoytopLevelUnionEmptyReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoytopLevelUnionEmpty& item);
bool BenchmarksVrutoytopLevelUnionEmptyWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoytopLevelUnionEmpty& item);

}} // namespace tl2::details

