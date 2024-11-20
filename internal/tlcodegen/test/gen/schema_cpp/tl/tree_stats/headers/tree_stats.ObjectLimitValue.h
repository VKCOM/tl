#pragma once

#include "../../../basics/basictl.h"
#include "../types/tree_stats.ObjectLimitValue.h"

namespace tl2 { namespace details { 

void TreeStatsObjectLimitValueReset(::tl2::tree_stats::ObjectLimitValue& item);

bool TreeStatsObjectLimitValueWriteJSON(std::ostream & s, const ::tl2::tree_stats::ObjectLimitValue& item);
bool TreeStatsObjectLimitValueReadBoxed(::basictl::tl_istream & s, ::tl2::tree_stats::ObjectLimitValue& item);
bool TreeStatsObjectLimitValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tree_stats::ObjectLimitValue& item);

}} // namespace tl2::details

