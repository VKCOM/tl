#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/integer.h"

namespace tl2 { namespace details { 

void BuiltinVectorIntegerReset(std::vector<::tl2::Integer>& item);

bool BuiltinVectorIntegerWriteJSON(std::ostream & s, const std::vector<::tl2::Integer>& item);
bool BuiltinVectorIntegerRead(::basictl::tl_istream & s, std::vector<::tl2::Integer>& item);
bool BuiltinVectorIntegerWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Integer>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void IntegerReset(::tl2::Integer& item) noexcept;

bool IntegerWriteJSON(std::ostream& s, const ::tl2::Integer& item) noexcept;
bool IntegerRead(::basictl::tl_istream & s, ::tl2::Integer& item) noexcept; 
bool IntegerWrite(::basictl::tl_ostream & s, const ::tl2::Integer& item) noexcept;
bool IntegerReadBoxed(::basictl::tl_istream & s, ::tl2::Integer& item);
bool IntegerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Integer& item);

}} // namespace tl2::details

