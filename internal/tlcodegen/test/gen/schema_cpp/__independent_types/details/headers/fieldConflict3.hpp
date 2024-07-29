#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../../__common/types/fieldConflict3.hpp"

namespace tl2 { namespace details { 

void FieldConflict3Reset(::tl2::FieldConflict3& item);
bool FieldConflict3Read(::basictl::tl_istream & s, ::tl2::FieldConflict3& item);
bool FieldConflict3Write(::basictl::tl_ostream & s, const ::tl2::FieldConflict3& item);
bool FieldConflict3ReadBoxed(::basictl::tl_istream & s, ::tl2::FieldConflict3& item);
bool FieldConflict3WriteBoxed(::basictl::tl_ostream & s, const ::tl2::FieldConflict3& item);

}} // namespace tl2::details

