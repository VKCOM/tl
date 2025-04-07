#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/service1.addOrGet.h"
#include "../types/service1.Value.h"

namespace tl2 { namespace details { 

void Service1AddOrGetReset(::tl2::service1::AddOrGet& item);

bool Service1AddOrGetWriteJSON(std::ostream& s, const ::tl2::service1::AddOrGet& item);
bool Service1AddOrGetRead(::basictl::tl_istream & s, ::tl2::service1::AddOrGet& item);
bool Service1AddOrGetWrite(::basictl::tl_ostream & s, const ::tl2::service1::AddOrGet& item);
bool Service1AddOrGetReadBoxed(::basictl::tl_istream & s, ::tl2::service1::AddOrGet& item);
bool Service1AddOrGetWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::AddOrGet& item);

bool Service1AddOrGetReadResult(::basictl::tl_istream & s, ::tl2::service1::AddOrGet& item, ::tl2::service1::Value& result);
bool Service1AddOrGetWriteResult(::basictl::tl_ostream & s, ::tl2::service1::AddOrGet& item, ::tl2::service1::Value& result);
		
}} // namespace tl2::details

