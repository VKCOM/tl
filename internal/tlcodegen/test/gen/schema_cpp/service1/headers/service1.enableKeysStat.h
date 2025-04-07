#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/service1.enableKeysStat.h"
#include "../../__common_namespace/types/Bool.h"

namespace tl2 { namespace details { 

void Service1EnableKeysStatReset(::tl2::service1::EnableKeysStat& item);

bool Service1EnableKeysStatWriteJSON(std::ostream& s, const ::tl2::service1::EnableKeysStat& item);
bool Service1EnableKeysStatRead(::basictl::tl_istream & s, ::tl2::service1::EnableKeysStat& item);
bool Service1EnableKeysStatWrite(::basictl::tl_ostream & s, const ::tl2::service1::EnableKeysStat& item);
bool Service1EnableKeysStatReadBoxed(::basictl::tl_istream & s, ::tl2::service1::EnableKeysStat& item);
bool Service1EnableKeysStatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::EnableKeysStat& item);

bool Service1EnableKeysStatReadResult(::basictl::tl_istream & s, ::tl2::service1::EnableKeysStat& item, bool& result);
bool Service1EnableKeysStatWriteResult(::basictl::tl_ostream & s, ::tl2::service1::EnableKeysStat& item, bool& result);
		
}} // namespace tl2::details

