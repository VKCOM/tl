#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/myDictOfInt.hpp"

namespace tl2 { namespace details { 

void MyDictOfIntReset(::tl2::MyDictOfInt& item);

bool MyDictOfIntWriteJSON(std::ostream& s, const ::tl2::MyDictOfInt& item);
bool MyDictOfIntRead(::basictl::tl_istream & s, ::tl2::MyDictOfInt& item);
bool MyDictOfIntWrite(::basictl::tl_ostream & s, const ::tl2::MyDictOfInt& item);
bool MyDictOfIntReadBoxed(::basictl::tl_istream & s, ::tl2::MyDictOfInt& item);
bool MyDictOfIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyDictOfInt& item);

}} // namespace tl2::details

