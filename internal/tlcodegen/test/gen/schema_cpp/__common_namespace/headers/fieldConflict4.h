// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/fieldConflict4.h"

namespace tlgen { namespace details { 

void FieldConflict4Reset(::tlgen::FieldConflict4& item) noexcept;

bool FieldConflict4WriteJSON(std::ostream& s, const ::tlgen::FieldConflict4& item) noexcept;
bool FieldConflict4Read(::tlgen::basictl::tl_istream & s, ::tlgen::FieldConflict4& item) noexcept; 
bool FieldConflict4Write(::tlgen::basictl::tl_ostream & s, const ::tlgen::FieldConflict4& item) noexcept;
bool FieldConflict4ReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::FieldConflict4& item);
bool FieldConflict4WriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::FieldConflict4& item);

}} // namespace tlgen::details

