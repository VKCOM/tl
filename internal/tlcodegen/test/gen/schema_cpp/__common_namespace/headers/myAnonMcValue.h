#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/myAnonMcValue.h"

namespace tl2 { namespace details { 

void MyAnonMcValueReset(::tl2::MyAnonMcValue& item) noexcept;

bool MyAnonMcValueWriteJSON(std::ostream& s, const ::tl2::MyAnonMcValue& item) noexcept;
bool MyAnonMcValueRead(::basictl::tl_istream & s, ::tl2::MyAnonMcValue& item) noexcept; 
bool MyAnonMcValueWrite(::basictl::tl_ostream & s, const ::tl2::MyAnonMcValue& item) noexcept;
bool MyAnonMcValueReadBoxed(::basictl::tl_istream & s, ::tl2::MyAnonMcValue& item);
bool MyAnonMcValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyAnonMcValue& item);

}} // namespace tl2::details

