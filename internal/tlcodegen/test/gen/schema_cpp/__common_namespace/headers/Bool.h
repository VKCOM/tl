// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/Bool.h"

namespace tlgen { namespace details { 

bool BoolWriteJSON(std::ostream & s, bool item);
bool BoolReadBoxed(::tlgen::basictl::tl_istream & s, bool& item);
bool BoolWriteBoxed(::tlgen::basictl::tl_ostream & s, bool item);

}} // namespace tlgen::details

