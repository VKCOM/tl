#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/myMcValue.h"

namespace tl2 { namespace details { 

void MyMcValueReset(::tl2::MyMcValue& item) noexcept;

bool MyMcValueWriteJSON(std::ostream& s, const ::tl2::MyMcValue& item) noexcept;
bool MyMcValueRead(::basictl::tl_istream & s, ::tl2::MyMcValue& item) noexcept; 
bool MyMcValueWrite(::basictl::tl_ostream & s, const ::tl2::MyMcValue& item) noexcept;
bool MyMcValueReadBoxed(::basictl::tl_istream & s, ::tl2::MyMcValue& item);
bool MyMcValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyMcValue& item);

}} // namespace tl2::details

