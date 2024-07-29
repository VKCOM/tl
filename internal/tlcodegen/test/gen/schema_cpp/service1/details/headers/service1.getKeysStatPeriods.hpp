#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../functions/service1.getKeysStatPeriods.hpp"
#include "../../../__common/types/int.hpp"

namespace tl2 { namespace details { 

void Service1GetKeysStatPeriodsReset(::tl2::service1::GetKeysStatPeriods& item);
bool Service1GetKeysStatPeriodsRead(::basictl::tl_istream & s, ::tl2::service1::GetKeysStatPeriods& item);
bool Service1GetKeysStatPeriodsWrite(::basictl::tl_ostream & s, const ::tl2::service1::GetKeysStatPeriods& item);
bool Service1GetKeysStatPeriodsReadBoxed(::basictl::tl_istream & s, ::tl2::service1::GetKeysStatPeriods& item);
bool Service1GetKeysStatPeriodsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::GetKeysStatPeriods& item);

bool Service1GetKeysStatPeriodsReadResult(::basictl::tl_istream & s, ::tl2::service1::GetKeysStatPeriods& item, std::vector<int32_t>& result);
bool Service1GetKeysStatPeriodsWriteResult(::basictl::tl_ostream & s, ::tl2::service1::GetKeysStatPeriods& item, std::vector<int32_t>& result);
		
}} // namespace tl2::details

