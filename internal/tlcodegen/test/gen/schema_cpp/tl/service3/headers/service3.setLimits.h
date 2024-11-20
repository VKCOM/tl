#pragma once

#include "../../../basics/basictl.h"
#include "../functions/service3.setLimits.h"
#include "../../__common_namespace/types/boolStat.h"

namespace tl2 { namespace details { 

void Service3SetLimitsReset(::tl2::service3::SetLimits& item);

bool Service3SetLimitsWriteJSON(std::ostream& s, const ::tl2::service3::SetLimits& item);
bool Service3SetLimitsRead(::basictl::tl_istream & s, ::tl2::service3::SetLimits& item);
bool Service3SetLimitsWrite(::basictl::tl_ostream & s, const ::tl2::service3::SetLimits& item);
bool Service3SetLimitsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::SetLimits& item);
bool Service3SetLimitsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::SetLimits& item);

bool Service3SetLimitsReadResult(::basictl::tl_istream & s, ::tl2::service3::SetLimits& item, ::tl2::BoolStat& result);
bool Service3SetLimitsWriteResult(::basictl::tl_ostream & s, ::tl2::service3::SetLimits& item, ::tl2::BoolStat& result);
		
}} // namespace tl2::details

