#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service1/functions/service1.getWildcardDict.h"
#include "__common_namespace/types/dictionaryField.h"

namespace tl2 { namespace details { 

void Service1GetWildcardDictReset(::tl2::service1::GetWildcardDict& item) noexcept;

bool Service1GetWildcardDictWriteJSON(std::ostream& s, const ::tl2::service1::GetWildcardDict& item) noexcept;
bool Service1GetWildcardDictRead(::basictl::tl_istream & s, ::tl2::service1::GetWildcardDict& item) noexcept; 
bool Service1GetWildcardDictWrite(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcardDict& item) noexcept;
bool Service1GetWildcardDictReadBoxed(::basictl::tl_istream & s, ::tl2::service1::GetWildcardDict& item);
bool Service1GetWildcardDictWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcardDict& item);

bool Service1GetWildcardDictReadResult(::basictl::tl_istream & s, ::tl2::service1::GetWildcardDict& item, std::map<std::string, std::string>& result);
bool Service1GetWildcardDictWriteResult(::basictl::tl_ostream & s, ::tl2::service1::GetWildcardDict& item, std::map<std::string, std::string>& result);
		
}} // namespace tl2::details

