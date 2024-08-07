#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/service1.getExpireTime.hpp"
#include "../../__common_namespace/types/int.hpp"

namespace tl2 { namespace details { 

void Service1GetExpireTimeReset(::tl2::service1::GetExpireTime& item);

bool Service1GetExpireTimeWriteJSON(std::ostream& s, const ::tl2::service1::GetExpireTime& item);
bool Service1GetExpireTimeRead(::basictl::tl_istream & s, ::tl2::service1::GetExpireTime& item);
bool Service1GetExpireTimeWrite(::basictl::tl_ostream & s, const ::tl2::service1::GetExpireTime& item);
bool Service1GetExpireTimeReadBoxed(::basictl::tl_istream & s, ::tl2::service1::GetExpireTime& item);
bool Service1GetExpireTimeWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::GetExpireTime& item);

bool Service1GetExpireTimeReadResult(::basictl::tl_istream & s, ::tl2::service1::GetExpireTime& item, std::optional<int32_t>& result);
bool Service1GetExpireTimeWriteResult(::basictl::tl_ostream & s, ::tl2::service1::GetExpireTime& item, std::optional<int32_t>& result);
		
}} // namespace tl2::details

