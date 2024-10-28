#pragma once

#include "../../../basics/basictl.h"
#include "../functions/service1.touch.h"
#include "../../__common_namespace/types/Bool.h"

namespace tl2 { namespace details { 

void Service1TouchReset(::tl2::service1::Touch& item);

bool Service1TouchWriteJSON(std::ostream& s, const ::tl2::service1::Touch& item);
bool Service1TouchRead(::basictl::tl_istream & s, ::tl2::service1::Touch& item);
bool Service1TouchWrite(::basictl::tl_ostream & s, const ::tl2::service1::Touch& item);
bool Service1TouchReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Touch& item);
bool Service1TouchWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Touch& item);

bool Service1TouchReadResult(::basictl::tl_istream & s, ::tl2::service1::Touch& item, bool& result);
bool Service1TouchWriteResult(::basictl::tl_ostream & s, ::tl2::service1::Touch& item, bool& result);
		
}} // namespace tl2::details
