// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/myNat.h"

namespace tl2 { namespace details { 

void MyNatReset(::tl2::MyNat& item) noexcept;

bool MyNatWriteJSON(std::ostream& s, const ::tl2::MyNat& item) noexcept;
bool MyNatRead(::basictl::tl_istream & s, ::tl2::MyNat& item) noexcept; 
bool MyNatWrite(::basictl::tl_ostream & s, const ::tl2::MyNat& item) noexcept;
bool MyNatReadBoxed(::basictl::tl_istream & s, ::tl2::MyNat& item);
bool MyNatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyNat& item);

}} // namespace tl2::details

