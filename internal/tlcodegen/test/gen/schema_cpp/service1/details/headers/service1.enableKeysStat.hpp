#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../functions/service1.enableKeysStat.hpp"
#include "../../../__common/types/Bool.hpp"

namespace tl2 { namespace details { 

void Service1EnableKeysStatReset(::tl2::service1::EnableKeysStat& item);
bool Service1EnableKeysStatRead(::basictl::tl_istream & s, ::tl2::service1::EnableKeysStat& item);
bool Service1EnableKeysStatWrite(::basictl::tl_ostream & s, const ::tl2::service1::EnableKeysStat& item);
bool Service1EnableKeysStatReadBoxed(::basictl::tl_istream & s, ::tl2::service1::EnableKeysStat& item);
bool Service1EnableKeysStatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::EnableKeysStat& item);

bool Service1EnableKeysStatReadResult(::basictl::tl_istream & s, ::tl2::service1::EnableKeysStat& item, bool& result);
bool Service1EnableKeysStatWriteResult(::basictl::tl_ostream & s, ::tl2::service1::EnableKeysStat& item, bool& result);
		
}} // namespace tl2::details

