#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/service1.setOrIncr.hpp"
#include "../types/service1.Value.hpp"

namespace tl2 { namespace details { 

void Service1SetOrIncrReset(::tl2::service1::SetOrIncr& item);

bool Service1SetOrIncrWriteJSON(std::ostream& s, const ::tl2::service1::SetOrIncr& item);
bool Service1SetOrIncrRead(::basictl::tl_istream & s, ::tl2::service1::SetOrIncr& item);
bool Service1SetOrIncrWrite(::basictl::tl_ostream & s, const ::tl2::service1::SetOrIncr& item);
bool Service1SetOrIncrReadBoxed(::basictl::tl_istream & s, ::tl2::service1::SetOrIncr& item);
bool Service1SetOrIncrWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::SetOrIncr& item);

bool Service1SetOrIncrReadResult(::basictl::tl_istream & s, ::tl2::service1::SetOrIncr& item, ::tl2::service1::Value& result);
bool Service1SetOrIncrWriteResult(::basictl::tl_ostream & s, ::tl2::service1::SetOrIncr& item, ::tl2::service1::Value& result);
		
}} // namespace tl2::details

