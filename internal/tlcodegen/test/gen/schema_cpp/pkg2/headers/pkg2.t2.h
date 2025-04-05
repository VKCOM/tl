#pragma once

#include "../../basictl/io_streams.h"
#include "../types/pkg2.t2.h"

namespace tl2 { namespace details { 

void Pkg2T2Reset(::tl2::pkg2::T2& item);

bool Pkg2T2WriteJSON(std::ostream& s, const ::tl2::pkg2::T2& item);
bool Pkg2T2Read(::basictl::tl_istream & s, ::tl2::pkg2::T2& item);
bool Pkg2T2Write(::basictl::tl_ostream & s, const ::tl2::pkg2::T2& item);
bool Pkg2T2ReadBoxed(::basictl::tl_istream & s, ::tl2::pkg2::T2& item);
bool Pkg2T2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::pkg2::T2& item);

}} // namespace tl2::details

