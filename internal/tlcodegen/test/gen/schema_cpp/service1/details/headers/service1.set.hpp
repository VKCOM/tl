#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../functions/service1.set.hpp"
#include "../../../__common/types/Bool.hpp"

namespace tl2 { namespace details { 

void Service1SetReset(::tl2::service1::Set& item);
bool Service1SetRead(::basictl::tl_istream & s, ::tl2::service1::Set& item);
bool Service1SetWrite(::basictl::tl_ostream & s, const ::tl2::service1::Set& item);
bool Service1SetReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Set& item);
bool Service1SetWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Set& item);

bool Service1SetReadResult(::basictl::tl_istream & s, ::tl2::service1::Set& item, bool& result);
bool Service1SetWriteResult(::basictl::tl_ostream & s, ::tl2::service1::Set& item, bool& result);
		
}} // namespace tl2::details

