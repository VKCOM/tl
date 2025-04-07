#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/get_arrays.h"
#include "../types/int.h"

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

