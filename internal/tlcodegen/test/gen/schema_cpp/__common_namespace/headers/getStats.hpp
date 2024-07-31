#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/getStats.hpp"
#include "../../tasks/types/tasks.queueTypeStats.hpp"

namespace tl2 { namespace details { 

void GetStatsReset(::tl2::GetStats& item);

bool GetStatsWriteJSON(std::ostream& s, const ::tl2::GetStats& item);
bool GetStatsRead(::basictl::tl_istream & s, ::tl2::GetStats& item);
bool GetStatsWrite(::basictl::tl_ostream & s, const ::tl2::GetStats& item);
bool GetStatsReadBoxed(::basictl::tl_istream & s, ::tl2::GetStats& item);
bool GetStatsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetStats& item);

bool GetStatsReadResult(::basictl::tl_istream & s, ::tl2::GetStats& item, ::tl2::tasks::QueueTypeStats& result);
bool GetStatsWriteResult(::basictl::tl_ostream & s, ::tl2::GetStats& item, ::tl2::tasks::QueueTypeStats& result);
		
}} // namespace tl2::details

