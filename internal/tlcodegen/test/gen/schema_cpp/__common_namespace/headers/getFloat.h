#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/functions/getFloat.h"

namespace tl2 { namespace details { 

void GetFloatReset(::tl2::GetFloat& item) noexcept;

bool GetFloatWriteJSON(std::ostream& s, const ::tl2::GetFloat& item) noexcept;
bool GetFloatRead(::basictl::tl_istream & s, ::tl2::GetFloat& item) noexcept; 
bool GetFloatWrite(::basictl::tl_ostream & s, const ::tl2::GetFloat& item) noexcept;
bool GetFloatReadBoxed(::basictl::tl_istream & s, ::tl2::GetFloat& item);
bool GetFloatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetFloat& item);

bool GetFloatReadResult(::basictl::tl_istream & s, ::tl2::GetFloat& item, float& result);
bool GetFloatWriteResult(::basictl::tl_ostream & s, ::tl2::GetFloat& item, float& result);
		
}} // namespace tl2::details

