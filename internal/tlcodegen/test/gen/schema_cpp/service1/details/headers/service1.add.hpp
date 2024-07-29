#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../functions/service1.add.hpp"
#include "../../../__common/types/Bool.hpp"

namespace tl2 { namespace details { 

void Service1AddReset(::tl2::service1::Add& item);
bool Service1AddRead(::basictl::tl_istream & s, ::tl2::service1::Add& item);
bool Service1AddWrite(::basictl::tl_ostream & s, const ::tl2::service1::Add& item);
bool Service1AddReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Add& item);
bool Service1AddWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Add& item);

bool Service1AddReadResult(::basictl::tl_istream & s, ::tl2::service1::Add& item, bool& result);
bool Service1AddWriteResult(::basictl::tl_ostream & s, ::tl2::service1::Add& item, bool& result);
		
}} // namespace tl2::details

