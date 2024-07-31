#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/getDouble.hpp"

namespace tl2 { namespace details { 

void GetDoubleReset(::tl2::GetDouble& item);

bool GetDoubleWriteJSON(std::ostream& s, const ::tl2::GetDouble& item);
bool GetDoubleRead(::basictl::tl_istream & s, ::tl2::GetDouble& item);
bool GetDoubleWrite(::basictl::tl_ostream & s, const ::tl2::GetDouble& item);
bool GetDoubleReadBoxed(::basictl::tl_istream & s, ::tl2::GetDouble& item);
bool GetDoubleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetDouble& item);

bool GetDoubleReadResult(::basictl::tl_istream & s, ::tl2::GetDouble& item, double& result);
bool GetDoubleWriteResult(::basictl::tl_ostream & s, ::tl2::GetDouble& item, double& result);
		
}} // namespace tl2::details

