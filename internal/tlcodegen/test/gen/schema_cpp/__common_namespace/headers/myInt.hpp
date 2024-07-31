#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/myInt.hpp"

namespace tl2 { namespace details { 

void MyIntReset(::tl2::MyInt& item);

bool MyIntWriteJSON(std::ostream& s, const ::tl2::MyInt& item);
bool MyIntRead(::basictl::tl_istream & s, ::tl2::MyInt& item);
bool MyIntWrite(::basictl::tl_ostream & s, const ::tl2::MyInt& item);
bool MyIntReadBoxed(::basictl::tl_istream & s, ::tl2::MyInt& item);
bool MyIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyInt& item);

}} // namespace tl2::details

