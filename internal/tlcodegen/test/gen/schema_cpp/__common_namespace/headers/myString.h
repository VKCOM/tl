// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/myString.h"

namespace tl2 { namespace details { 

void MyStringReset(::tl2::MyString& item) noexcept;

bool MyStringWriteJSON(std::ostream& s, const ::tl2::MyString& item) noexcept;
bool MyStringRead(::basictl::tl_istream & s, ::tl2::MyString& item) noexcept; 
bool MyStringWrite(::basictl::tl_ostream & s, const ::tl2::MyString& item) noexcept;
bool MyStringReadBoxed(::basictl::tl_istream & s, ::tl2::MyString& item);
bool MyStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyString& item);

}} // namespace tl2::details

