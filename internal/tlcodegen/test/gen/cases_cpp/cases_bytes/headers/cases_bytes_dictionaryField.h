#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../../__common_namespace/types/dictionaryField.h"

namespace tl2 { namespace details { 

void BuiltinVectorDictionaryFieldStringReset(std::map<std::string, std::string>& item);

bool BuiltinVectorDictionaryFieldStringWriteJSON(std::ostream & s, const std::map<std::string, std::string>& item);
bool BuiltinVectorDictionaryFieldStringRead(::basictl::tl_istream & s, std::map<std::string, std::string>& item);
bool BuiltinVectorDictionaryFieldStringWrite(::basictl::tl_ostream & s, const std::map<std::string, std::string>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryFieldStringReset(::tl2::DictionaryField<std::string>& item) noexcept;

bool DictionaryFieldStringWriteJSON(std::ostream& s, const ::tl2::DictionaryField<std::string>& item) noexcept;
bool DictionaryFieldStringRead(::basictl::tl_istream & s, ::tl2::DictionaryField<std::string>& item) noexcept; 
bool DictionaryFieldStringWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryField<std::string>& item) noexcept;
bool DictionaryFieldStringReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryField<std::string>& item);
bool DictionaryFieldStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryField<std::string>& item);

}} // namespace tl2::details

