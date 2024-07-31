#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/tree_stats.objectLimitValueDouble.hpp"

namespace tl2 { namespace details { 

void TreeStatsObjectLimitValueDoubleReset(::tl2::tree_stats::ObjectLimitValueDouble& item);

bool TreeStatsObjectLimitValueDoubleWriteJSON(std::ostream& s, const ::tl2::tree_stats::ObjectLimitValueDouble& item);
bool TreeStatsObjectLimitValueDoubleRead(::basictl::tl_istream & s, ::tl2::tree_stats::ObjectLimitValueDouble& item);
bool TreeStatsObjectLimitValueDoubleWrite(::basictl::tl_ostream & s, const ::tl2::tree_stats::ObjectLimitValueDouble& item);
bool TreeStatsObjectLimitValueDoubleReadBoxed(::basictl::tl_istream & s, ::tl2::tree_stats::ObjectLimitValueDouble& item);
bool TreeStatsObjectLimitValueDoubleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tree_stats::ObjectLimitValueDouble& item);

}} // namespace tl2::details

