#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/vector.h"
#include "__common_namespace/types/dictionaryField.h"

namespace tl2 { namespace details { 

void VectorDictionaryFieldDictionaryIntReset(std::map<std::string, std::map<std::string, int32_t>>& item) noexcept;

bool VectorDictionaryFieldDictionaryIntWriteJSON(std::ostream& s, const std::map<std::string, std::map<std::string, int32_t>>& item) noexcept;
bool VectorDictionaryFieldDictionaryIntRead(::basictl::tl_istream & s, std::map<std::string, std::map<std::string, int32_t>>& item) noexcept; 
bool VectorDictionaryFieldDictionaryIntWrite(::basictl::tl_ostream & s, const std::map<std::string, std::map<std::string, int32_t>>& item) noexcept;
bool VectorDictionaryFieldDictionaryIntReadBoxed(::basictl::tl_istream & s, std::map<std::string, std::map<std::string, int32_t>>& item);
bool VectorDictionaryFieldDictionaryIntWriteBoxed(::basictl::tl_ostream & s, const std::map<std::string, std::map<std::string, int32_t>>& item);

}} // namespace tl2::details

