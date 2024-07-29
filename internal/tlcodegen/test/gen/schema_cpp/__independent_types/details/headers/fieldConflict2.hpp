#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../../__common/types/fieldConflict2.hpp"

namespace tl2 { namespace details { 

void FieldConflict2Reset(::tl2::FieldConflict2& item);
bool FieldConflict2Read(::basictl::tl_istream & s, ::tl2::FieldConflict2& item);
bool FieldConflict2Write(::basictl::tl_ostream & s, const ::tl2::FieldConflict2& item);
bool FieldConflict2ReadBoxed(::basictl::tl_istream & s, ::tl2::FieldConflict2& item);
bool FieldConflict2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::FieldConflict2& item);

}} // namespace tl2::details

