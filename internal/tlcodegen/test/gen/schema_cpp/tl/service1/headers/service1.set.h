#pragma once

#include "../../../basics/basictl.h"
#include "../functions/service1.set.h"
#include "../../__common_namespace/types/Bool.h"

namespace tl2 { namespace details { 

void Service1SetReset(::tl2::service1::Set& item);

bool Service1SetWriteJSON(std::ostream& s, const ::tl2::service1::Set& item);
bool Service1SetRead(::basictl::tl_istream & s, ::tl2::service1::Set& item);
bool Service1SetWrite(::basictl::tl_ostream & s, const ::tl2::service1::Set& item);
bool Service1SetReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Set& item);
bool Service1SetWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Set& item);

bool Service1SetReadResult(::basictl::tl_istream & s, ::tl2::service1::Set& item, bool& result);
bool Service1SetWriteResult(::basictl::tl_ostream & s, ::tl2::service1::Set& item, bool& result);
		
}} // namespace tl2::details

