// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cd/types/cd.myType.h"

namespace tl2 { namespace details { 

void CdMyTypeReset(::tl2::cd::MyType& item) noexcept;

bool CdMyTypeWriteJSON(std::ostream& s, const ::tl2::cd::MyType& item) noexcept;
bool CdMyTypeRead(::basictl::tl_istream & s, ::tl2::cd::MyType& item) noexcept; 
bool CdMyTypeWrite(::basictl::tl_ostream & s, const ::tl2::cd::MyType& item) noexcept;
bool CdMyTypeReadBoxed(::basictl::tl_istream & s, ::tl2::cd::MyType& item);
bool CdMyTypeWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cd::MyType& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

bool CdMyTypeMaybeWriteJSON(std::ostream & s, const std::optional<::tl2::cd::MyType>& item);

bool CdMyTypeMaybeReadBoxed(::basictl::tl_istream & s, std::optional<::tl2::cd::MyType>& item);
bool CdMyTypeMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<::tl2::cd::MyType>& item);


}} // namespace tl2::details

