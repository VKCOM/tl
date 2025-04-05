#pragma once

#include "../../basictl/io_streams.h"
#include "../types/int32.h"

namespace tl2 { namespace details { 

void Int32Reset(::tl2::Int32& item);

bool Int32WriteJSON(std::ostream& s, const ::tl2::Int32& item);
bool Int32Read(::basictl::tl_istream & s, ::tl2::Int32& item);
bool Int32Write(::basictl::tl_ostream & s, const ::tl2::Int32& item);
bool Int32ReadBoxed(::basictl::tl_istream & s, ::tl2::Int32& item);
bool Int32WriteBoxed(::basictl::tl_ostream & s, const ::tl2::Int32& item);

}} // namespace tl2::details

