#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/service3.getLastVisitTimestamp.h"
#include "../../__common_namespace/types/int.h"

namespace tl2 { namespace details { 

void Service3GetLastVisitTimestampReset(::tl2::service3::GetLastVisitTimestamp& item);

bool Service3GetLastVisitTimestampWriteJSON(std::ostream& s, const ::tl2::service3::GetLastVisitTimestamp& item);
bool Service3GetLastVisitTimestampRead(::basictl::tl_istream & s, ::tl2::service3::GetLastVisitTimestamp& item);
bool Service3GetLastVisitTimestampWrite(::basictl::tl_ostream & s, const ::tl2::service3::GetLastVisitTimestamp& item);
bool Service3GetLastVisitTimestampReadBoxed(::basictl::tl_istream & s, ::tl2::service3::GetLastVisitTimestamp& item);
bool Service3GetLastVisitTimestampWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::GetLastVisitTimestamp& item);

bool Service3GetLastVisitTimestampReadResult(::basictl::tl_istream & s, ::tl2::service3::GetLastVisitTimestamp& item, std::optional<int32_t>& result);
bool Service3GetLastVisitTimestampWriteResult(::basictl::tl_ostream & s, ::tl2::service3::GetLastVisitTimestamp& item, std::optional<int32_t>& result);
		
}} // namespace tl2::details

