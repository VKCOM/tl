// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cd/types/cd.myType.h"

namespace tlgen { namespace details { 

bool CdMyTypeMaybeWriteJSON(std::ostream & s, const std::optional<::tlgen::cd::MyType>& item);

bool CdMyTypeMaybeReadBoxed(::tlgen::basictl::tl_istream & s, std::optional<::tlgen::cd::MyType>& item);
bool CdMyTypeMaybeWriteBoxed(::tlgen::basictl::tl_ostream & s, const std::optional<::tlgen::cd::MyType>& item);


}} // namespace tlgen::details

