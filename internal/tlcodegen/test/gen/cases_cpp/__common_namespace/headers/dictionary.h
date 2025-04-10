#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/dictionary.h"
#include "../types/dictionaryField.h"

namespace tl2 { namespace details { 

void DictionaryIntReset(std::map<std::string, int32_t>& item);

bool DictionaryIntWriteJSON(std::ostream& s, const std::map<std::string, int32_t>& item);
bool DictionaryIntRead(::basictl::tl_istream & s, std::map<std::string, int32_t>& item);
bool DictionaryIntWrite(::basictl::tl_ostream & s, const std::map<std::string, int32_t>& item);
bool DictionaryIntReadBoxed(::basictl::tl_istream & s, std::map<std::string, int32_t>& item);
bool DictionaryIntWriteBoxed(::basictl::tl_ostream & s, const std::map<std::string, int32_t>& item);

}} // namespace tl2::details

