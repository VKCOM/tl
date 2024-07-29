#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../functions/service3.getLastVisitTimestamp.hpp"
#include "../../../__common/types/int.hpp"

namespace tl2 { namespace details { 

void Service3GetLastVisitTimestampReset(::tl2::service3::GetLastVisitTimestamp& item);
bool Service3GetLastVisitTimestampRead(::basictl::tl_istream & s, ::tl2::service3::GetLastVisitTimestamp& item);
bool Service3GetLastVisitTimestampWrite(::basictl::tl_ostream & s, const ::tl2::service3::GetLastVisitTimestamp& item);
bool Service3GetLastVisitTimestampReadBoxed(::basictl::tl_istream & s, ::tl2::service3::GetLastVisitTimestamp& item);
bool Service3GetLastVisitTimestampWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::GetLastVisitTimestamp& item);

bool Service3GetLastVisitTimestampReadResult(::basictl::tl_istream & s, ::tl2::service3::GetLastVisitTimestamp& item, std::optional<int32_t>& result);
bool Service3GetLastVisitTimestampWriteResult(::basictl::tl_ostream & s, ::tl2::service3::GetLastVisitTimestamp& item, std::optional<int32_t>& result);
		
}} // namespace tl2::details

