#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../functions/service1.disableKeysStat.hpp"
#include "../../../__common/types/Bool.hpp"

namespace tl2 { namespace details { 

void Service1DisableKeysStatReset(::tl2::service1::DisableKeysStat& item);
bool Service1DisableKeysStatRead(::basictl::tl_istream & s, ::tl2::service1::DisableKeysStat& item);
bool Service1DisableKeysStatWrite(::basictl::tl_ostream & s, const ::tl2::service1::DisableKeysStat& item);
bool Service1DisableKeysStatReadBoxed(::basictl::tl_istream & s, ::tl2::service1::DisableKeysStat& item);
bool Service1DisableKeysStatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::DisableKeysStat& item);

bool Service1DisableKeysStatReadResult(::basictl::tl_istream & s, ::tl2::service1::DisableKeysStat& item, bool& result);
bool Service1DisableKeysStatWriteResult(::basictl::tl_ostream & s, ::tl2::service1::DisableKeysStat& item, bool& result);
		
}} // namespace tl2::details

