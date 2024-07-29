#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../functions/service3.getLimits.hpp"
#include "../../types/service3.limits.hpp"

namespace tl2 { namespace details { 

void Service3GetLimitsReset(::tl2::service3::GetLimits& item);
bool Service3GetLimitsRead(::basictl::tl_istream & s, ::tl2::service3::GetLimits& item);
bool Service3GetLimitsWrite(::basictl::tl_ostream & s, const ::tl2::service3::GetLimits& item);
bool Service3GetLimitsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::GetLimits& item);
bool Service3GetLimitsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::GetLimits& item);

bool Service3GetLimitsReadResult(::basictl::tl_istream & s, ::tl2::service3::GetLimits& item, ::tl2::service3::Limits& result);
bool Service3GetLimitsWriteResult(::basictl::tl_ostream & s, ::tl2::service3::GetLimits& item, ::tl2::service3::Limits& result);
		
}} // namespace tl2::details

