#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/service1.get.hpp"
#include "../types/service1.Value.hpp"

namespace tl2 { namespace details { 

void Service1GetReset(::tl2::service1::Get& item);
bool Service1GetRead(::basictl::tl_istream & s, ::tl2::service1::Get& item);
bool Service1GetWrite(::basictl::tl_ostream & s, const ::tl2::service1::Get& item);
bool Service1GetReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Get& item);
bool Service1GetWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Get& item);

bool Service1GetReadResult(::basictl::tl_istream & s, ::tl2::service1::Get& item, ::tl2::service1::Value& result);
bool Service1GetWriteResult(::basictl::tl_ostream & s, ::tl2::service1::Get& item, ::tl2::service1::Value& result);
		
}} // namespace tl2::details

