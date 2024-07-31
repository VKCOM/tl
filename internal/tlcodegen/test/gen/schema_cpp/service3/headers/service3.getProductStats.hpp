#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/service3.getProductStats.hpp"
#include "../../__common_namespace/types/vector.hpp"
#include "../types/service3.productStatsOld.hpp"

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

