#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/tree_stats.ObjectLimitValue.h"

namespace tl2 { namespace details { 

void TreeStatsObjectLimitValueReset(::tl2::tree_stats::ObjectLimitValue& item) noexcept;

bool TreeStatsObjectLimitValueWriteJSON(std::ostream & s, const ::tl2::tree_stats::ObjectLimitValue& item) noexcept;
bool TreeStatsObjectLimitValueReadBoxed(::basictl::tl_istream & s, ::tl2::tree_stats::ObjectLimitValue& item) noexcept;
bool TreeStatsObjectLimitValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tree_stats::ObjectLimitValue& item) noexcept;

}} // namespace tl2::details

