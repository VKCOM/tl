#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/integer.hpp"

namespace tl2 { namespace details { 

void BuiltinVectorIntegerReset(std::vector<::tl2::Integer>& item);
bool BuiltinVectorIntegerRead(::basictl::tl_istream & s, std::vector<::tl2::Integer>& item);
bool BuiltinVectorIntegerWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Integer>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void IntegerReset(::tl2::Integer& item);
bool IntegerRead(::basictl::tl_istream & s, ::tl2::Integer& item);
bool IntegerWrite(::basictl::tl_ostream & s, const ::tl2::Integer& item);
bool IntegerReadBoxed(::basictl::tl_istream & s, ::tl2::Integer& item);
bool IntegerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Integer& item);

}} // namespace tl2::details

