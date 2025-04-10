#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/service1.getWildcardWithFlags.h"
#include "../types/service1.Value.h"
#include "../../__common_namespace/types/dictionaryField.h"

namespace tl2 { namespace details { 

void Service1GetWildcardWithFlagsReset(::tl2::service1::GetWildcardWithFlags& item);

bool Service1GetWildcardWithFlagsWriteJSON(std::ostream& s, const ::tl2::service1::GetWildcardWithFlags& item);
bool Service1GetWildcardWithFlagsRead(::basictl::tl_istream & s, ::tl2::service1::GetWildcardWithFlags& item);
bool Service1GetWildcardWithFlagsWrite(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcardWithFlags& item);
bool Service1GetWildcardWithFlagsReadBoxed(::basictl::tl_istream & s, ::tl2::service1::GetWildcardWithFlags& item);
bool Service1GetWildcardWithFlagsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcardWithFlags& item);

bool Service1GetWildcardWithFlagsReadResult(::basictl::tl_istream & s, ::tl2::service1::GetWildcardWithFlags& item, std::map<std::string, ::tl2::service1::Value>& result);
bool Service1GetWildcardWithFlagsWriteResult(::basictl::tl_ostream & s, ::tl2::service1::GetWildcardWithFlags& item, std::map<std::string, ::tl2::service1::Value>& result);
		
}} // namespace tl2::details

