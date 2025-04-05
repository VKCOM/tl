#pragma once

#include "../../basictl/io_streams.h"
#include "../types/dictionaryField.h"

namespace tl2 { namespace details { 

void BuiltinVectorDictionaryFieldIntReset(std::vector<::tl2::DictionaryField<int32_t>>& item);

bool BuiltinVectorDictionaryFieldIntWriteJSON(std::ostream & s, const std::vector<::tl2::DictionaryField<int32_t>>& item);
bool BuiltinVectorDictionaryFieldIntRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryField<int32_t>>& item);
bool BuiltinVectorDictionaryFieldIntWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryField<int32_t>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryFieldIntReset(::tl2::DictionaryField<int32_t>& item);

bool DictionaryFieldIntWriteJSON(std::ostream& s, const ::tl2::DictionaryField<int32_t>& item);
bool DictionaryFieldIntRead(::basictl::tl_istream & s, ::tl2::DictionaryField<int32_t>& item);
bool DictionaryFieldIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryField<int32_t>& item);
bool DictionaryFieldIntReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryField<int32_t>& item);
bool DictionaryFieldIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryField<int32_t>& item);

}} // namespace tl2::details

