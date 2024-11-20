#pragma once

#include "../../../basics/basictl.h"
#include "../functions/service1.getWildcardDict.h"
#include "../../__common_namespace/types/dictionary.h"

namespace tl2 { namespace details { 

void Service1GetWildcardDictReset(::tl2::service1::GetWildcardDict& item);

bool Service1GetWildcardDictWriteJSON(std::ostream& s, const ::tl2::service1::GetWildcardDict& item);
bool Service1GetWildcardDictRead(::basictl::tl_istream & s, ::tl2::service1::GetWildcardDict& item);
bool Service1GetWildcardDictWrite(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcardDict& item);
bool Service1GetWildcardDictReadBoxed(::basictl::tl_istream & s, ::tl2::service1::GetWildcardDict& item);
bool Service1GetWildcardDictWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcardDict& item);

bool Service1GetWildcardDictReadResult(::basictl::tl_istream & s, ::tl2::service1::GetWildcardDict& item, ::tl2::Dictionary<std::string>& result);
bool Service1GetWildcardDictWriteResult(::basictl::tl_ostream & s, ::tl2::service1::GetWildcardDict& item, ::tl2::Dictionary<std::string>& result);
		
}} // namespace tl2::details

