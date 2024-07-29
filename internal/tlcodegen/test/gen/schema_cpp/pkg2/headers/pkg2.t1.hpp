#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/pkg2.t1.hpp"

namespace tl2 { namespace details { 

void Pkg2T1Reset(::tl2::pkg2::T1& item);
bool Pkg2T1Read(::basictl::tl_istream & s, ::tl2::pkg2::T1& item);
bool Pkg2T1Write(::basictl::tl_ostream & s, const ::tl2::pkg2::T1& item);
bool Pkg2T1ReadBoxed(::basictl::tl_istream & s, ::tl2::pkg2::T1& item);
bool Pkg2T1WriteBoxed(::basictl::tl_ostream & s, const ::tl2::pkg2::T1& item);

}} // namespace tl2::details

