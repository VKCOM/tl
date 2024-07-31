#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/myMcValue.hpp"

namespace tl2 { namespace details { 

void MyMcValueReset(::tl2::MyMcValue& item);

bool MyMcValueWriteJSON(std::ostream& s, const ::tl2::MyMcValue& item);
bool MyMcValueRead(::basictl::tl_istream & s, ::tl2::MyMcValue& item);
bool MyMcValueWrite(::basictl::tl_ostream & s, const ::tl2::MyMcValue& item);
bool MyMcValueReadBoxed(::basictl::tl_istream & s, ::tl2::MyMcValue& item);
bool MyMcValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyMcValue& item);

}} // namespace tl2::details

