#pragma once

#include "../../../basics/basictl.h"
#include "../types/tree_stats.objectLimitValueLong.h"

namespace tl2 { namespace details { 

void TreeStatsObjectLimitValueLongReset(::tl2::tree_stats::ObjectLimitValueLong& item);

bool TreeStatsObjectLimitValueLongWriteJSON(std::ostream& s, const ::tl2::tree_stats::ObjectLimitValueLong& item);
bool TreeStatsObjectLimitValueLongRead(::basictl::tl_istream & s, ::tl2::tree_stats::ObjectLimitValueLong& item);
bool TreeStatsObjectLimitValueLongWrite(::basictl::tl_ostream & s, const ::tl2::tree_stats::ObjectLimitValueLong& item);
bool TreeStatsObjectLimitValueLongReadBoxed(::basictl::tl_istream & s, ::tl2::tree_stats::ObjectLimitValueLong& item);
bool TreeStatsObjectLimitValueLongWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tree_stats::ObjectLimitValueLong& item);

}} // namespace tl2::details

