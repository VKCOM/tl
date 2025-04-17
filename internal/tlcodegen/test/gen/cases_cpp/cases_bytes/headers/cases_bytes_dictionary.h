#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/dictionary.h"
#include "__common_namespace/types/dictionaryField.h"

namespace tl2 { namespace details { 

void DictionaryStringReset(std::map<std::string, std::string>& item) noexcept;

bool DictionaryStringWriteJSON(std::ostream& s, const std::map<std::string, std::string>& item) noexcept;
bool DictionaryStringRead(::basictl::tl_istream & s, std::map<std::string, std::string>& item) noexcept; 
bool DictionaryStringWrite(::basictl::tl_ostream & s, const std::map<std::string, std::string>& item) noexcept;
bool DictionaryStringReadBoxed(::basictl::tl_istream & s, std::map<std::string, std::string>& item);
bool DictionaryStringWriteBoxed(::basictl::tl_ostream & s, const std::map<std::string, std::string>& item);

}} // namespace tl2::details

