#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/myString.h"

namespace tl2 { namespace details { 

void MyStringReset(::tl2::MyString& item);

bool MyStringWriteJSON(std::ostream& s, const ::tl2::MyString& item);
bool MyStringRead(::basictl::tl_istream & s, ::tl2::MyString& item);
bool MyStringWrite(::basictl::tl_ostream & s, const ::tl2::MyString& item);
bool MyStringReadBoxed(::basictl::tl_istream & s, ::tl2::MyString& item);
bool MyStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyString& item);

}} // namespace tl2::details

