#pragma once

#include "../../../basics/basictl.h"
#include "../functions/service1.addOrIncr.h"
#include "../types/service1.Value.h"

namespace tl2 { namespace details { 

void Service1AddOrIncrReset(::tl2::service1::AddOrIncr& item);

bool Service1AddOrIncrWriteJSON(std::ostream& s, const ::tl2::service1::AddOrIncr& item);
bool Service1AddOrIncrRead(::basictl::tl_istream & s, ::tl2::service1::AddOrIncr& item);
bool Service1AddOrIncrWrite(::basictl::tl_ostream & s, const ::tl2::service1::AddOrIncr& item);
bool Service1AddOrIncrReadBoxed(::basictl::tl_istream & s, ::tl2::service1::AddOrIncr& item);
bool Service1AddOrIncrWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::AddOrIncr& item);

bool Service1AddOrIncrReadResult(::basictl::tl_istream & s, ::tl2::service1::AddOrIncr& item, ::tl2::service1::Value& result);
bool Service1AddOrIncrWriteResult(::basictl::tl_ostream & s, ::tl2::service1::AddOrIncr& item, ::tl2::service1::Value& result);
		
}} // namespace tl2::details

