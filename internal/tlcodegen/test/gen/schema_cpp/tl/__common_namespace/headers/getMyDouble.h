#pragma once

#include "../../../basics/basictl.h"
#include "../functions/getMyDouble.h"
#include "../types/myDouble.h"

namespace tl2 { namespace details { 

void GetMyDoubleReset(::tl2::GetMyDouble& item);

bool GetMyDoubleWriteJSON(std::ostream& s, const ::tl2::GetMyDouble& item);
bool GetMyDoubleRead(::basictl::tl_istream & s, ::tl2::GetMyDouble& item);
bool GetMyDoubleWrite(::basictl::tl_ostream & s, const ::tl2::GetMyDouble& item);
bool GetMyDoubleReadBoxed(::basictl::tl_istream & s, ::tl2::GetMyDouble& item);
bool GetMyDoubleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetMyDouble& item);

bool GetMyDoubleReadResult(::basictl::tl_istream & s, ::tl2::GetMyDouble& item, ::tl2::MyDouble& result);
bool GetMyDoubleWriteResult(::basictl::tl_ostream & s, ::tl2::GetMyDouble& item, ::tl2::MyDouble& result);
		
}} // namespace tl2::details

