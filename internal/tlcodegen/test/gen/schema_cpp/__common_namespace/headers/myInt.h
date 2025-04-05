#pragma once

#include "../../basictl/io_streams.h"
#include "../types/myInt.h"

namespace tl2 { namespace details { 

void MyIntReset(::tl2::MyInt& item);

bool MyIntWriteJSON(std::ostream& s, const ::tl2::MyInt& item);
bool MyIntRead(::basictl::tl_istream & s, ::tl2::MyInt& item);
bool MyIntWrite(::basictl::tl_ostream & s, const ::tl2::MyInt& item);
bool MyIntReadBoxed(::basictl::tl_istream & s, ::tl2::MyInt& item);
bool MyIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyInt& item);

}} // namespace tl2::details

