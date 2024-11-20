#pragma once

#include "../../../basics/basictl.h"
#include "../types/fieldConflict1.h"

namespace tl2 { namespace details { 

void FieldConflict1Reset(::tl2::FieldConflict1& item);

bool FieldConflict1WriteJSON(std::ostream& s, const ::tl2::FieldConflict1& item);
bool FieldConflict1Read(::basictl::tl_istream & s, ::tl2::FieldConflict1& item);
bool FieldConflict1Write(::basictl::tl_ostream & s, const ::tl2::FieldConflict1& item);
bool FieldConflict1ReadBoxed(::basictl::tl_istream & s, ::tl2::FieldConflict1& item);
bool FieldConflict1WriteBoxed(::basictl::tl_ostream & s, const ::tl2::FieldConflict1& item);

}} // namespace tl2::details

