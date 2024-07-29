#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/service1.getWildcardWithFlags.hpp"
#include "../types/service1.Value.hpp"
#include "../../__common_namespace/types/dictionary.hpp"

namespace tl2 { namespace details { 

void Service1GetWildcardWithFlagsReset(::tl2::service1::GetWildcardWithFlags& item);
bool Service1GetWildcardWithFlagsRead(::basictl::tl_istream & s, ::tl2::service1::GetWildcardWithFlags& item);
bool Service1GetWildcardWithFlagsWrite(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcardWithFlags& item);
bool Service1GetWildcardWithFlagsReadBoxed(::basictl::tl_istream & s, ::tl2::service1::GetWildcardWithFlags& item);
bool Service1GetWildcardWithFlagsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcardWithFlags& item);

bool Service1GetWildcardWithFlagsReadResult(::basictl::tl_istream & s, ::tl2::service1::GetWildcardWithFlags& item, ::tl2::Dictionary<::tl2::service1::Value>& result);
bool Service1GetWildcardWithFlagsWriteResult(::basictl::tl_ostream & s, ::tl2::service1::GetWildcardWithFlags& item, ::tl2::Dictionary<::tl2::service1::Value>& result);
		
}} // namespace tl2::details

