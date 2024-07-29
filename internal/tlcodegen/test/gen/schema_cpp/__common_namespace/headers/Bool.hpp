#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/Bool.hpp"

namespace tl2 { namespace details { 

bool BoolReadBoxed(::basictl::tl_istream & s, bool& item);
bool BoolWriteBoxed(::basictl::tl_ostream & s, bool item);

}} // namespace tl2::details

