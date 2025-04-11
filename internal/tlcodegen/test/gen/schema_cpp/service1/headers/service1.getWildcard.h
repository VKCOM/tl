#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/service1.getWildcard.h"
#include "../../__common_namespace/types/map.h"

namespace tl2 { namespace details { 

void Service1GetWildcardReset(::tl2::service1::GetWildcard& item) noexcept;

bool Service1GetWildcardWriteJSON(std::ostream& s, const ::tl2::service1::GetWildcard& item) noexcept;
bool Service1GetWildcardRead(::basictl::tl_istream & s, ::tl2::service1::GetWildcard& item) noexcept; 
bool Service1GetWildcardWrite(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcard& item) noexcept;
bool Service1GetWildcardReadBoxed(::basictl::tl_istream & s, ::tl2::service1::GetWildcard& item);
bool Service1GetWildcardWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcard& item);

bool Service1GetWildcardReadResult(::basictl::tl_istream & s, ::tl2::service1::GetWildcard& item, std::vector<::tl2::Map<std::string, std::string>>& result);
bool Service1GetWildcardWriteResult(::basictl::tl_ostream & s, ::tl2::service1::GetWildcard& item, std::vector<::tl2::Map<std::string, std::string>>& result);
		
}} // namespace tl2::details

