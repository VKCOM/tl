#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/fieldConflict3.h"

namespace tl2 { namespace details { 

void FieldConflict3Reset(::tl2::FieldConflict3& item) noexcept;

bool FieldConflict3WriteJSON(std::ostream& s, const ::tl2::FieldConflict3& item) noexcept;
bool FieldConflict3Read(::basictl::tl_istream & s, ::tl2::FieldConflict3& item) noexcept; 
bool FieldConflict3Write(::basictl::tl_ostream & s, const ::tl2::FieldConflict3& item) noexcept;
bool FieldConflict3ReadBoxed(::basictl::tl_istream & s, ::tl2::FieldConflict3& item);
bool FieldConflict3WriteBoxed(::basictl::tl_ostream & s, const ::tl2::FieldConflict3& item);

}} // namespace tl2::details

