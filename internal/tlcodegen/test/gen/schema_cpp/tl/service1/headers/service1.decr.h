#pragma once

#include "../../../basics/basictl.h"
#include "../functions/service1.decr.h"
#include "../types/service1.Value.h"

namespace tl2 { namespace details { 

void Service1DecrReset(::tl2::service1::Decr& item);

bool Service1DecrWriteJSON(std::ostream& s, const ::tl2::service1::Decr& item);
bool Service1DecrRead(::basictl::tl_istream & s, ::tl2::service1::Decr& item);
bool Service1DecrWrite(::basictl::tl_ostream & s, const ::tl2::service1::Decr& item);
bool Service1DecrReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Decr& item);
bool Service1DecrWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Decr& item);

bool Service1DecrReadResult(::basictl::tl_istream & s, ::tl2::service1::Decr& item, ::tl2::service1::Value& result);
bool Service1DecrWriteResult(::basictl::tl_ostream & s, ::tl2::service1::Decr& item, ::tl2::service1::Value& result);
		
}} // namespace tl2::details

