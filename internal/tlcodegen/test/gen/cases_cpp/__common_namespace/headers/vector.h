#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/vector.h"
#include "__common_namespace/types/dictionaryField.h"

namespace tl2 { namespace details { 

void VectorDictionaryFieldIntReset(std::map<std::string, int32_t>& item) noexcept;

bool VectorDictionaryFieldIntWriteJSON(std::ostream& s, const std::map<std::string, int32_t>& item) noexcept;
bool VectorDictionaryFieldIntRead(::basictl::tl_istream & s, std::map<std::string, int32_t>& item) noexcept; 
bool VectorDictionaryFieldIntWrite(::basictl::tl_ostream & s, const std::map<std::string, int32_t>& item) noexcept;
bool VectorDictionaryFieldIntReadBoxed(::basictl::tl_istream & s, std::map<std::string, int32_t>& item);
bool VectorDictionaryFieldIntWriteBoxed(::basictl::tl_ostream & s, const std::map<std::string, int32_t>& item);

}} // namespace tl2::details

