#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../../__common/types/vector.hpp"
#include "../../../__common/types/dictionaryField.hpp"
#include "../../../__common/types/string.hpp"

namespace tl2 { namespace details { 

void VectorDictionaryFieldStringReset(std::vector<::tl2::DictionaryField<std::string>>& item);
bool VectorDictionaryFieldStringRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryField<std::string>>& item);
bool VectorDictionaryFieldStringWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryField<std::string>>& item);
bool VectorDictionaryFieldStringReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::DictionaryField<std::string>>& item);
bool VectorDictionaryFieldStringWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryField<std::string>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorStringReset(std::vector<std::string>& item);
bool VectorStringRead(::basictl::tl_istream & s, std::vector<std::string>& item);
bool VectorStringWrite(::basictl::tl_ostream & s, const std::vector<std::string>& item);
bool VectorStringReadBoxed(::basictl::tl_istream & s, std::vector<std::string>& item);
bool VectorStringWriteBoxed(::basictl::tl_ostream & s, const std::vector<std::string>& item);

}} // namespace tl2::details

