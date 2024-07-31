#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/myString.hpp"

namespace tl2 { namespace details { 

void MyStringReset(::tl2::MyString& item);

bool MyStringWriteJSON(std::ostream& s, const ::tl2::MyString& item);
bool MyStringRead(::basictl::tl_istream & s, ::tl2::MyString& item);
bool MyStringWrite(::basictl::tl_ostream & s, const ::tl2::MyString& item);
bool MyStringReadBoxed(::basictl::tl_istream & s, ::tl2::MyString& item);
bool MyStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyString& item);

}} // namespace tl2::details

