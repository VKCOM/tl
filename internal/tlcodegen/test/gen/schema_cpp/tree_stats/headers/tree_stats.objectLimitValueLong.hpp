#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/tree_stats.objectLimitValueLong.hpp"

namespace tl2 { namespace details { 

void TreeStatsObjectLimitValueLongReset(::tl2::tree_stats::ObjectLimitValueLong& item);

bool TreeStatsObjectLimitValueLongWriteJSON(std::ostream& s, const ::tl2::tree_stats::ObjectLimitValueLong& item);
bool TreeStatsObjectLimitValueLongRead(::basictl::tl_istream & s, ::tl2::tree_stats::ObjectLimitValueLong& item);
bool TreeStatsObjectLimitValueLongWrite(::basictl::tl_ostream & s, const ::tl2::tree_stats::ObjectLimitValueLong& item);
bool TreeStatsObjectLimitValueLongReadBoxed(::basictl::tl_istream & s, ::tl2::tree_stats::ObjectLimitValueLong& item);
bool TreeStatsObjectLimitValueLongWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tree_stats::ObjectLimitValueLong& item);

}} // namespace tl2::details

