#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../../__common_namespace/types/dictionary.h"
#include "../../__common_namespace/types/dictionaryField.h"

namespace tl2 { namespace details { 

void DictionaryDictionaryIntReset(std::map<std::string, std::map<std::string, int32_t>>& item);

bool DictionaryDictionaryIntWriteJSON(std::ostream& s, const std::map<std::string, std::map<std::string, int32_t>>& item);
bool DictionaryDictionaryIntRead(::basictl::tl_istream & s, std::map<std::string, std::map<std::string, int32_t>>& item);
bool DictionaryDictionaryIntWrite(::basictl::tl_ostream & s, const std::map<std::string, std::map<std::string, int32_t>>& item);
bool DictionaryDictionaryIntReadBoxed(::basictl::tl_istream & s, std::map<std::string, std::map<std::string, int32_t>>& item);
bool DictionaryDictionaryIntWriteBoxed(::basictl::tl_ostream & s, const std::map<std::string, std::map<std::string, int32_t>>& item);

}} // namespace tl2::details

