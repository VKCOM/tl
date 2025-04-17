#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "tree_stats/types/tree_stats.objectLimitValueLong.h"

namespace tl2 { namespace details { 

void TreeStatsObjectLimitValueLongReset(::tl2::tree_stats::ObjectLimitValueLong& item) noexcept;

bool TreeStatsObjectLimitValueLongWriteJSON(std::ostream& s, const ::tl2::tree_stats::ObjectLimitValueLong& item) noexcept;
bool TreeStatsObjectLimitValueLongRead(::basictl::tl_istream & s, ::tl2::tree_stats::ObjectLimitValueLong& item) noexcept; 
bool TreeStatsObjectLimitValueLongWrite(::basictl::tl_ostream & s, const ::tl2::tree_stats::ObjectLimitValueLong& item) noexcept;
bool TreeStatsObjectLimitValueLongReadBoxed(::basictl::tl_istream & s, ::tl2::tree_stats::ObjectLimitValueLong& item);
bool TreeStatsObjectLimitValueLongWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tree_stats::ObjectLimitValueLong& item);

}} // namespace tl2::details

