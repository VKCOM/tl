#pragma once

#include "../../basictl/io_streams.h"
#include "../../__common_namespace/types/dictionaryField.h"

namespace tl2 { namespace details { 

void BuiltinVectorDictionaryFieldStringReset(std::vector<::tl2::DictionaryField<std::string>>& item);

bool BuiltinVectorDictionaryFieldStringWriteJSON(std::ostream & s, const std::vector<::tl2::DictionaryField<std::string>>& item);
bool BuiltinVectorDictionaryFieldStringRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryField<std::string>>& item);
bool BuiltinVectorDictionaryFieldStringWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryField<std::string>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryFieldStringReset(::tl2::DictionaryField<std::string>& item);

bool DictionaryFieldStringWriteJSON(std::ostream& s, const ::tl2::DictionaryField<std::string>& item);
bool DictionaryFieldStringRead(::basictl::tl_istream & s, ::tl2::DictionaryField<std::string>& item);
bool DictionaryFieldStringWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryField<std::string>& item);
bool DictionaryFieldStringReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryField<std::string>& item);
bool DictionaryFieldStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryField<std::string>& item);

}} // namespace tl2::details

