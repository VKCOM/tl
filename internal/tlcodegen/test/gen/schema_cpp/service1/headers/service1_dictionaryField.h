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

namespace tl2 { namespace details { 

void DictionaryFieldDictionaryIntReset(::tl2::DictionaryField<std::map<std::string, int32_t>>& item);

bool DictionaryFieldDictionaryIntWriteJSON(std::ostream& s, const ::tl2::DictionaryField<std::map<std::string, int32_t>>& item);
bool DictionaryFieldDictionaryIntRead(::basictl::tl_istream & s, ::tl2::DictionaryField<std::map<std::string, int32_t>>& item);
bool DictionaryFieldDictionaryIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryField<std::map<std::string, int32_t>>& item);
bool DictionaryFieldDictionaryIntReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryField<std::map<std::string, int32_t>>& item);
bool DictionaryFieldDictionaryIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryField<std::map<std::string, int32_t>>& item);

}} // namespace tl2::details

