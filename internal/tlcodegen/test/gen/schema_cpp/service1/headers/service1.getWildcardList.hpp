#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/service1.getWildcardList.hpp"
#include "../../__common_namespace/types/string.hpp"

namespace tl2 { namespace details { 

void Service1GetWildcardListReset(::tl2::service1::GetWildcardList& item);

bool Service1GetWildcardListWriteJSON(std::ostream& s, const ::tl2::service1::GetWildcardList& item);
bool Service1GetWildcardListRead(::basictl::tl_istream & s, ::tl2::service1::GetWildcardList& item);
bool Service1GetWildcardListWrite(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcardList& item);
bool Service1GetWildcardListReadBoxed(::basictl::tl_istream & s, ::tl2::service1::GetWildcardList& item);
bool Service1GetWildcardListWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcardList& item);

bool Service1GetWildcardListReadResult(::basictl::tl_istream & s, ::tl2::service1::GetWildcardList& item, std::vector<std::string>& result);
bool Service1GetWildcardListWriteResult(::basictl::tl_ostream & s, ::tl2::service1::GetWildcardList& item, std::vector<std::string>& result);
		
}} // namespace tl2::details

