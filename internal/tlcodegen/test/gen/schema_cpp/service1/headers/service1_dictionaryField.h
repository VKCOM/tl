#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../../__common_namespace/types/dictionaryField.h"

namespace tl2 { namespace details { 

void BuiltinVectorDictionaryFieldDictionaryIntReset(std::map<std::string, std::map<std::string, int32_t>>& item);

bool BuiltinVectorDictionaryFieldDictionaryIntWriteJSON(std::ostream & s, const std::map<std::string, std::map<std::string, int32_t>>& item);
bool BuiltinVectorDictionaryFieldDictionaryIntRead(::basictl::tl_istream & s, std::map<std::string, std::map<std::string, int32_t>>& item);
bool BuiltinVectorDictionaryFieldDictionaryIntWrite(::basictl::tl_ostream & s, const std::map<std::string, std::map<std::string, int32_t>>& item);

}} // namespace tl2::details

