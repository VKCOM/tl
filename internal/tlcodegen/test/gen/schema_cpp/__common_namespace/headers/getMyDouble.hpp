#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/getMyDouble.hpp"
#include "../types/myDouble.hpp"

namespace tl2 { namespace details { 

void GetMyDoubleReset(::tl2::GetMyDouble& item);
bool GetMyDoubleRead(::basictl::tl_istream & s, ::tl2::GetMyDouble& item);
bool GetMyDoubleWrite(::basictl::tl_ostream & s, const ::tl2::GetMyDouble& item);
bool GetMyDoubleReadBoxed(::basictl::tl_istream & s, ::tl2::GetMyDouble& item);
bool GetMyDoubleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetMyDouble& item);

bool GetMyDoubleReadResult(::basictl::tl_istream & s, ::tl2::GetMyDouble& item, ::tl2::MyDouble& result);
bool GetMyDoubleWriteResult(::basictl::tl_ostream & s, ::tl2::GetMyDouble& item, ::tl2::MyDouble& result);
		
}} // namespace tl2::details

