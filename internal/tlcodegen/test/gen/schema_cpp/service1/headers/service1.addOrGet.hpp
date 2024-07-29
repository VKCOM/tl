#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/service1.addOrGet.hpp"
#include "../types/service1.Value.hpp"

namespace tl2 { namespace details { 

void Service1AddOrGetReset(::tl2::service1::AddOrGet& item);
bool Service1AddOrGetRead(::basictl::tl_istream & s, ::tl2::service1::AddOrGet& item);
bool Service1AddOrGetWrite(::basictl::tl_ostream & s, const ::tl2::service1::AddOrGet& item);
bool Service1AddOrGetReadBoxed(::basictl::tl_istream & s, ::tl2::service1::AddOrGet& item);
bool Service1AddOrGetWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::AddOrGet& item);

bool Service1AddOrGetReadResult(::basictl::tl_istream & s, ::tl2::service1::AddOrGet& item, ::tl2::service1::Value& result);
bool Service1AddOrGetWriteResult(::basictl::tl_ostream & s, ::tl2::service1::AddOrGet& item, ::tl2::service1::Value& result);
		
}} // namespace tl2::details

