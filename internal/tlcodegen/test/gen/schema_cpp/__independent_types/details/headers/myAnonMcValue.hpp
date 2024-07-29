#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../../__common/types/myAnonMcValue.hpp"

namespace tl2 { namespace details { 

void MyAnonMcValueReset(::tl2::MyAnonMcValue& item);
bool MyAnonMcValueRead(::basictl::tl_istream & s, ::tl2::MyAnonMcValue& item);
bool MyAnonMcValueWrite(::basictl::tl_ostream & s, const ::tl2::MyAnonMcValue& item);
bool MyAnonMcValueReadBoxed(::basictl::tl_istream & s, ::tl2::MyAnonMcValue& item);
bool MyAnonMcValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyAnonMcValue& item);

}} // namespace tl2::details

