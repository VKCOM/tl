#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/benchmarks.vrutoytopLevelUnionEmpty.h"

namespace tl2 { namespace details { 

void BenchmarksVrutoytopLevelUnionEmptyReset(::tl2::benchmarks::VrutoytopLevelUnionEmpty& item) noexcept;

bool BenchmarksVrutoytopLevelUnionEmptyWriteJSON(std::ostream& s, const ::tl2::benchmarks::VrutoytopLevelUnionEmpty& item) noexcept;
bool BenchmarksVrutoytopLevelUnionEmptyRead(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoytopLevelUnionEmpty& item) noexcept; 
bool BenchmarksVrutoytopLevelUnionEmptyWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoytopLevelUnionEmpty& item) noexcept;
bool BenchmarksVrutoytopLevelUnionEmptyReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoytopLevelUnionEmpty& item);
bool BenchmarksVrutoytopLevelUnionEmptyWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoytopLevelUnionEmpty& item);

}} // namespace tl2::details

