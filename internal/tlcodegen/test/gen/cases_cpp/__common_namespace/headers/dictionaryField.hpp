#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/dictionaryField.hpp"

namespace tl2 { namespace details { 

void BuiltinVectorDictionaryFieldIntReset(std::vector<::tl2::DictionaryField<int32_t>>& item);
bool BuiltinVectorDictionaryFieldIntRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryField<int32_t>>& item);
bool BuiltinVectorDictionaryFieldIntWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryField<int32_t>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryFieldIntReset(::tl2::DictionaryField<int32_t>& item);
bool DictionaryFieldIntRead(::basictl::tl_istream & s, ::tl2::DictionaryField<int32_t>& item);
bool DictionaryFieldIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryField<int32_t>& item);
bool DictionaryFieldIntReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryField<int32_t>& item);
bool DictionaryFieldIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryField<int32_t>& item);

}} // namespace tl2::details

