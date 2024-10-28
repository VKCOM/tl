#pragma once

#include "../../../basics/basictl.h"
#include "../functions/service3.getProductStats.h"
#include "../../__common_namespace/types/vector.h"
#include "../types/service3.productStatsOld.h"

namespace tl2 { namespace details { 

void Service3GetProductStatsReset(::tl2::service3::GetProductStats& item);

bool Service3GetProductStatsWriteJSON(std::ostream& s, const ::tl2::service3::GetProductStats& item);
bool Service3GetProductStatsRead(::basictl::tl_istream & s, ::tl2::service3::GetProductStats& item);
bool Service3GetProductStatsWrite(::basictl::tl_ostream & s, const ::tl2::service3::GetProductStats& item);
bool Service3GetProductStatsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::GetProductStats& item);
bool Service3GetProductStatsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::GetProductStats& item);

bool Service3GetProductStatsReadResult(::basictl::tl_istream & s, ::tl2::service3::GetProductStats& item, std::optional<std::vector<::tl2::service3::ProductStatsOld>>& result);
bool Service3GetProductStatsWriteResult(::basictl::tl_ostream & s, ::tl2::service3::GetProductStats& item, std::optional<std::vector<::tl2::service3::ProductStatsOld>>& result);
		
}} // namespace tl2::details
