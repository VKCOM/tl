#pragma once

#include "../../../basics/basictl.h"
#include "../types/Bool.h"

namespace tl2 { namespace details { 

bool BoolWriteJSON(std::ostream & s, bool item);
bool BoolReadBoxed(::basictl::tl_istream & s, bool& item);
bool BoolWriteBoxed(::basictl::tl_ostream & s, bool item);

}} // namespace tl2::details
