#pragma once

#include "../../../basics/basictl.h"
#include "../types/float.h"

namespace tl2 { namespace details { 

void FloatReset(float& item);

bool FloatWriteJSON(std::ostream& s, const float& item);
bool FloatRead(::basictl::tl_istream & s, float& item);
bool FloatWrite(::basictl::tl_ostream & s, const float& item);
bool FloatReadBoxed(::basictl::tl_istream & s, float& item);
bool FloatWriteBoxed(::basictl::tl_ostream & s, const float& item);

}} // namespace tl2::details

