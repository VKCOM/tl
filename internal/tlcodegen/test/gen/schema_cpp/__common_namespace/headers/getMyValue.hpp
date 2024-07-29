#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/getMyValue.hpp"
#include "../types/MyValue.hpp"

namespace tl2 { namespace details { 

void GetMyValueReset(::tl2::GetMyValue& item);
bool GetMyValueRead(::basictl::tl_istream & s, ::tl2::GetMyValue& item);
bool GetMyValueWrite(::basictl::tl_ostream & s, const ::tl2::GetMyValue& item);
bool GetMyValueReadBoxed(::basictl::tl_istream & s, ::tl2::GetMyValue& item);
bool GetMyValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetMyValue& item);

bool GetMyValueReadResult(::basictl::tl_istream & s, ::tl2::GetMyValue& item, ::tl2::MyValue& result);
bool GetMyValueWriteResult(::basictl::tl_ostream & s, ::tl2::GetMyValue& item, ::tl2::MyValue& result);
		
}} // namespace tl2::details

