#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/service1.incr.hpp"
#include "../types/service1.Value.hpp"

namespace tl2 { namespace details { 

void Service1IncrReset(::tl2::service1::Incr& item);
bool Service1IncrRead(::basictl::tl_istream & s, ::tl2::service1::Incr& item);
bool Service1IncrWrite(::basictl::tl_ostream & s, const ::tl2::service1::Incr& item);
bool Service1IncrReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Incr& item);
bool Service1IncrWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Incr& item);

bool Service1IncrReadResult(::basictl::tl_istream & s, ::tl2::service1::Incr& item, ::tl2::service1::Value& result);
bool Service1IncrWriteResult(::basictl::tl_ostream & s, ::tl2::service1::Incr& item, ::tl2::service1::Value& result);
		
}} // namespace tl2::details

