#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service3.productStatsOld.hpp"

namespace tl2 { namespace details { 

void BuiltinVectorService3ProductStatsOldReset(std::vector<::tl2::service3::ProductStatsOld>& item);
bool BuiltinVectorService3ProductStatsOldRead(::basictl::tl_istream & s, std::vector<::tl2::service3::ProductStatsOld>& item);
bool BuiltinVectorService3ProductStatsOldWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service3::ProductStatsOld>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void Service3ProductStatsOldReset(::tl2::service3::ProductStatsOld& item);
bool Service3ProductStatsOldRead(::basictl::tl_istream & s, ::tl2::service3::ProductStatsOld& item);
bool Service3ProductStatsOldWrite(::basictl::tl_ostream & s, const ::tl2::service3::ProductStatsOld& item);
bool Service3ProductStatsOldReadBoxed(::basictl::tl_istream & s, ::tl2::service3::ProductStatsOld& item);
bool Service3ProductStatsOldWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::ProductStatsOld& item);

}} // namespace tl2::details

