#pragma once

#include "../../basictl/io_streams.h"
#include "../types/vector.h"
#include "../types/dictionaryField.h"

namespace tl2 { namespace details { 

void VectorDictionaryFieldIntReset(std::vector<::tl2::DictionaryField<int32_t>>& item);

bool VectorDictionaryFieldIntWriteJSON(std::ostream& s, const std::vector<::tl2::DictionaryField<int32_t>>& item);
bool VectorDictionaryFieldIntRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryField<int32_t>>& item);
bool VectorDictionaryFieldIntWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryField<int32_t>>& item);
bool VectorDictionaryFieldIntReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::DictionaryField<int32_t>>& item);
bool VectorDictionaryFieldIntWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryField<int32_t>>& item);

}} // namespace tl2::details

