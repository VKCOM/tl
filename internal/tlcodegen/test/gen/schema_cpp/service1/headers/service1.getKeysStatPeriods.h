#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service1/functions/service1.getKeysStatPeriods.h"
#include "__common_namespace/types/int.h"

namespace tl2 { namespace details { 

void Service1GetKeysStatPeriodsReset(::tl2::service1::GetKeysStatPeriods& item) noexcept;

bool Service1GetKeysStatPeriodsWriteJSON(std::ostream& s, const ::tl2::service1::GetKeysStatPeriods& item) noexcept;
bool Service1GetKeysStatPeriodsRead(::basictl::tl_istream & s, ::tl2::service1::GetKeysStatPeriods& item) noexcept; 
bool Service1GetKeysStatPeriodsWrite(::basictl::tl_ostream & s, const ::tl2::service1::GetKeysStatPeriods& item) noexcept;
bool Service1GetKeysStatPeriodsReadBoxed(::basictl::tl_istream & s, ::tl2::service1::GetKeysStatPeriods& item);
bool Service1GetKeysStatPeriodsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::GetKeysStatPeriods& item);

bool Service1GetKeysStatPeriodsReadResult(::basictl::tl_istream & s, ::tl2::service1::GetKeysStatPeriods& item, std::vector<int32_t>& result);
bool Service1GetKeysStatPeriodsWriteResult(::basictl::tl_ostream & s, ::tl2::service1::GetKeysStatPeriods& item, std::vector<int32_t>& result);
		
}} // namespace tl2::details

