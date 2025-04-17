#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/myInt.h"

namespace tl2 { namespace details { 

void MyIntReset(::tl2::MyInt& item) noexcept;

bool MyIntWriteJSON(std::ostream& s, const ::tl2::MyInt& item) noexcept;
bool MyIntRead(::basictl::tl_istream & s, ::tl2::MyInt& item) noexcept; 
bool MyIntWrite(::basictl::tl_ostream & s, const ::tl2::MyInt& item) noexcept;
bool MyIntReadBoxed(::basictl::tl_istream & s, ::tl2::MyInt& item);
bool MyIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyInt& item);

}} // namespace tl2::details

