#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../../__common/types/int64.hpp"

namespace tl2 { namespace details { 

void Int64Reset(::tl2::Int64& item);
bool Int64Read(::basictl::tl_istream & s, ::tl2::Int64& item);
bool Int64Write(::basictl::tl_ostream & s, const ::tl2::Int64& item);
bool Int64ReadBoxed(::basictl::tl_istream & s, ::tl2::Int64& item);
bool Int64WriteBoxed(::basictl::tl_ostream & s, const ::tl2::Int64& item);

}} // namespace tl2::details

