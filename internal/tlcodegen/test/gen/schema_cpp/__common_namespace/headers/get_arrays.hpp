#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/get_arrays.hpp"
#include "../types/int.hpp"

namespace tl2 { namespace details { 

void GetArraysReset(::tl2::Get_arrays& item);

bool GetArraysWriteJSON(std::ostream& s, const ::tl2::Get_arrays& item);
bool GetArraysRead(::basictl::tl_istream & s, ::tl2::Get_arrays& item);
bool GetArraysWrite(::basictl::tl_ostream & s, const ::tl2::Get_arrays& item);
bool GetArraysReadBoxed(::basictl::tl_istream & s, ::tl2::Get_arrays& item);
bool GetArraysWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Get_arrays& item);

bool GetArraysReadResult(::basictl::tl_istream & s, ::tl2::Get_arrays& item, std::array<int32_t, 5>& result);
bool GetArraysWriteResult(::basictl::tl_ostream & s, ::tl2::Get_arrays& item, std::array<int32_t, 5>& result);
		
}} // namespace tl2::details

