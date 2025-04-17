#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/fieldConflict4.h"

namespace tl2 { namespace details { 

void FieldConflict4Reset(::tl2::FieldConflict4& item) noexcept;

bool FieldConflict4WriteJSON(std::ostream& s, const ::tl2::FieldConflict4& item) noexcept;
bool FieldConflict4Read(::basictl::tl_istream & s, ::tl2::FieldConflict4& item) noexcept; 
bool FieldConflict4Write(::basictl::tl_ostream & s, const ::tl2::FieldConflict4& item) noexcept;
bool FieldConflict4ReadBoxed(::basictl::tl_istream & s, ::tl2::FieldConflict4& item);
bool FieldConflict4WriteBoxed(::basictl::tl_ostream & s, const ::tl2::FieldConflict4& item);

}} // namespace tl2::details

