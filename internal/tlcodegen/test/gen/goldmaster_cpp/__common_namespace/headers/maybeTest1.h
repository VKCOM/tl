// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/maybeTest1.h"

namespace tl2 { namespace details { 

void MaybeTest1Reset(::tl2::MaybeTest1& item) noexcept;

bool MaybeTest1WriteJSON(std::ostream& s, const ::tl2::MaybeTest1& item) noexcept;
bool MaybeTest1Read(::basictl::tl_istream & s, ::tl2::MaybeTest1& item) noexcept; 
bool MaybeTest1Write(::basictl::tl_ostream & s, const ::tl2::MaybeTest1& item) noexcept;
bool MaybeTest1ReadBoxed(::basictl::tl_istream & s, ::tl2::MaybeTest1& item);
bool MaybeTest1WriteBoxed(::basictl::tl_ostream & s, const ::tl2::MaybeTest1& item);

}} // namespace tl2::details

