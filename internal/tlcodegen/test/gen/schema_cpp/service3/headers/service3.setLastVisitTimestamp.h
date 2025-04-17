#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service3/functions/service3.setLastVisitTimestamp.h"
#include "__common_namespace/types/Bool.h"

namespace tl2 { namespace details { 

void Service3SetLastVisitTimestampReset(::tl2::service3::SetLastVisitTimestamp& item) noexcept;

bool Service3SetLastVisitTimestampWriteJSON(std::ostream& s, const ::tl2::service3::SetLastVisitTimestamp& item) noexcept;
bool Service3SetLastVisitTimestampRead(::basictl::tl_istream & s, ::tl2::service3::SetLastVisitTimestamp& item) noexcept; 
bool Service3SetLastVisitTimestampWrite(::basictl::tl_ostream & s, const ::tl2::service3::SetLastVisitTimestamp& item) noexcept;
bool Service3SetLastVisitTimestampReadBoxed(::basictl::tl_istream & s, ::tl2::service3::SetLastVisitTimestamp& item);
bool Service3SetLastVisitTimestampWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::SetLastVisitTimestamp& item);

bool Service3SetLastVisitTimestampReadResult(::basictl::tl_istream & s, ::tl2::service3::SetLastVisitTimestamp& item, bool& result);
bool Service3SetLastVisitTimestampWriteResult(::basictl::tl_ostream & s, ::tl2::service3::SetLastVisitTimestamp& item, bool& result);
		
}} // namespace tl2::details

