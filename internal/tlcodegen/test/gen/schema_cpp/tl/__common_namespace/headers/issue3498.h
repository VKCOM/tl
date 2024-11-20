#pragma once

#include "../../../basics/basictl.h"
#include "../types/issue3498.h"

namespace tl2 { namespace details { 

void Issue3498Reset(::tl2::Issue3498& item);

bool Issue3498WriteJSON(std::ostream& s, const ::tl2::Issue3498& item);
bool Issue3498Read(::basictl::tl_istream & s, ::tl2::Issue3498& item);
bool Issue3498Write(::basictl::tl_ostream & s, const ::tl2::Issue3498& item);
bool Issue3498ReadBoxed(::basictl::tl_istream & s, ::tl2::Issue3498& item);
bool Issue3498WriteBoxed(::basictl::tl_ostream & s, const ::tl2::Issue3498& item);

}} // namespace tl2::details

