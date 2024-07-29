#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../functions/service1.getWildcard.hpp"
#include "../../../__common/types/map.hpp"

namespace tl2 { namespace details { 

void Service1GetWildcardReset(::tl2::service1::GetWildcard& item);
bool Service1GetWildcardRead(::basictl::tl_istream & s, ::tl2::service1::GetWildcard& item);
bool Service1GetWildcardWrite(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcard& item);
bool Service1GetWildcardReadBoxed(::basictl::tl_istream & s, ::tl2::service1::GetWildcard& item);
bool Service1GetWildcardWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcard& item);

bool Service1GetWildcardReadResult(::basictl::tl_istream & s, ::tl2::service1::GetWildcard& item, std::vector<::tl2::Map<std::string, std::string>>& result);
bool Service1GetWildcardWriteResult(::basictl::tl_ostream & s, ::tl2::service1::GetWildcard& item, std::vector<::tl2::Map<std::string, std::string>>& result);
		
}} // namespace tl2::details

