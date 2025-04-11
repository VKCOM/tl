#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/fieldConflict2.h"

namespace tl2 { namespace details { 

void FieldConflict2Reset(::tl2::FieldConflict2& item) noexcept;

bool FieldConflict2WriteJSON(std::ostream& s, const ::tl2::FieldConflict2& item) noexcept;
bool FieldConflict2Read(::basictl::tl_istream & s, ::tl2::FieldConflict2& item) noexcept; 
bool FieldConflict2Write(::basictl::tl_ostream & s, const ::tl2::FieldConflict2& item) noexcept;
bool FieldConflict2ReadBoxed(::basictl::tl_istream & s, ::tl2::FieldConflict2& item);
bool FieldConflict2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::FieldConflict2& item);

}} // namespace tl2::details

