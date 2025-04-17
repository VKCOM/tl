#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/functions/getDouble.h"

namespace tl2 { namespace details { 

void GetDoubleReset(::tl2::GetDouble& item) noexcept;

bool GetDoubleWriteJSON(std::ostream& s, const ::tl2::GetDouble& item) noexcept;
bool GetDoubleRead(::basictl::tl_istream & s, ::tl2::GetDouble& item) noexcept; 
bool GetDoubleWrite(::basictl::tl_ostream & s, const ::tl2::GetDouble& item) noexcept;
bool GetDoubleReadBoxed(::basictl::tl_istream & s, ::tl2::GetDouble& item);
bool GetDoubleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetDouble& item);

bool GetDoubleReadResult(::basictl::tl_istream & s, ::tl2::GetDouble& item, double& result);
bool GetDoubleWriteResult(::basictl::tl_ostream & s, ::tl2::GetDouble& item, double& result);
		
}} // namespace tl2::details

