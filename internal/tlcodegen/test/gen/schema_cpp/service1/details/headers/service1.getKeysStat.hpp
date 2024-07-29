#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../functions/service1.getKeysStat.hpp"
#include "../../types/service1.keysStat.hpp"

namespace tl2 { namespace details { 

void Service1GetKeysStatReset(::tl2::service1::GetKeysStat& item);
bool Service1GetKeysStatRead(::basictl::tl_istream & s, ::tl2::service1::GetKeysStat& item);
bool Service1GetKeysStatWrite(::basictl::tl_ostream & s, const ::tl2::service1::GetKeysStat& item);
bool Service1GetKeysStatReadBoxed(::basictl::tl_istream & s, ::tl2::service1::GetKeysStat& item);
bool Service1GetKeysStatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::GetKeysStat& item);

bool Service1GetKeysStatReadResult(::basictl::tl_istream & s, ::tl2::service1::GetKeysStat& item, std::optional<::tl2::service1::KeysStat>& result);
bool Service1GetKeysStatWriteResult(::basictl::tl_ostream & s, ::tl2::service1::GetKeysStat& item, std::optional<::tl2::service1::KeysStat>& result);
		
}} // namespace tl2::details

