#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/myDictOfInt.h"

namespace tl2 { namespace details { 

void MyDictOfIntReset(::tl2::MyDictOfInt& item) noexcept;

bool MyDictOfIntWriteJSON(std::ostream& s, const ::tl2::MyDictOfInt& item) noexcept;
bool MyDictOfIntRead(::basictl::tl_istream & s, ::tl2::MyDictOfInt& item) noexcept; 
bool MyDictOfIntWrite(::basictl::tl_ostream & s, const ::tl2::MyDictOfInt& item) noexcept;
bool MyDictOfIntReadBoxed(::basictl::tl_istream & s, ::tl2::MyDictOfInt& item);
bool MyDictOfIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyDictOfInt& item);

}} // namespace tl2::details

