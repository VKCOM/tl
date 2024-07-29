#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../../__common/types/fieldConflict1.hpp"

namespace tl2 { namespace details { 

void FieldConflict1Reset(::tl2::FieldConflict1& item);
bool FieldConflict1Read(::basictl::tl_istream & s, ::tl2::FieldConflict1& item);
bool FieldConflict1Write(::basictl::tl_ostream & s, const ::tl2::FieldConflict1& item);
bool FieldConflict1ReadBoxed(::basictl::tl_istream & s, ::tl2::FieldConflict1& item);
bool FieldConflict1WriteBoxed(::basictl::tl_ostream & s, const ::tl2::FieldConflict1& item);

}} // namespace tl2::details

