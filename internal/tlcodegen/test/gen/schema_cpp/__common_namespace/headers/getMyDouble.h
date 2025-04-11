#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/getMyDouble.h"
#include "../types/myDouble.h"

namespace tl2 { namespace details { 

void GetMyDoubleReset(::tl2::GetMyDouble& item) noexcept;

bool GetMyDoubleWriteJSON(std::ostream& s, const ::tl2::GetMyDouble& item) noexcept;
bool GetMyDoubleRead(::basictl::tl_istream & s, ::tl2::GetMyDouble& item) noexcept; 
bool GetMyDoubleWrite(::basictl::tl_ostream & s, const ::tl2::GetMyDouble& item) noexcept;
bool GetMyDoubleReadBoxed(::basictl::tl_istream & s, ::tl2::GetMyDouble& item);
bool GetMyDoubleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetMyDouble& item);

bool GetMyDoubleReadResult(::basictl::tl_istream & s, ::tl2::GetMyDouble& item, ::tl2::MyDouble& result);
bool GetMyDoubleWriteResult(::basictl::tl_ostream & s, ::tl2::GetMyDouble& item, ::tl2::MyDouble& result);
		
}} // namespace tl2::details

