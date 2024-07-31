#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../__common_namespace/types/vector.hpp"
#include "../types/service1.Value.hpp"
#include "../../__common_namespace/types/dictionary.hpp"
#include "../../__common_namespace/types/dictionaryField.hpp"

namespace tl2 { namespace details { 

void VectorDictionaryFieldDictionaryIntReset(std::vector<::tl2::DictionaryField<::tl2::Dictionary<int32_t>>>& item);

bool VectorDictionaryFieldDictionaryIntWriteJSON(std::ostream& s, const std::vector<::tl2::DictionaryField<::tl2::Dictionary<int32_t>>>& item);
bool VectorDictionaryFieldDictionaryIntRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryField<::tl2::Dictionary<int32_t>>>& item);
bool VectorDictionaryFieldDictionaryIntWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryField<::tl2::Dictionary<int32_t>>>& item);
bool VectorDictionaryFieldDictionaryIntReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::DictionaryField<::tl2::Dictionary<int32_t>>>& item);
bool VectorDictionaryFieldDictionaryIntWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryField<::tl2::Dictionary<int32_t>>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorDictionaryFieldService1ValueReset(std::vector<::tl2::DictionaryField<::tl2::service1::Value>>& item);

bool VectorDictionaryFieldService1ValueWriteJSON(std::ostream& s, const std::vector<::tl2::DictionaryField<::tl2::service1::Value>>& item);
bool VectorDictionaryFieldService1ValueRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryField<::tl2::service1::Value>>& item);
bool VectorDictionaryFieldService1ValueWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryField<::tl2::service1::Value>>& item);
bool VectorDictionaryFieldService1ValueReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::DictionaryField<::tl2::service1::Value>>& item);
bool VectorDictionaryFieldService1ValueWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryField<::tl2::service1::Value>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorDictionaryFieldStringReset(std::vector<::tl2::DictionaryField<std::string>>& item);

bool VectorDictionaryFieldStringWriteJSON(std::ostream& s, const std::vector<::tl2::DictionaryField<std::string>>& item);
bool VectorDictionaryFieldStringRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryField<std::string>>& item);
bool VectorDictionaryFieldStringWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryField<std::string>>& item);
bool VectorDictionaryFieldStringReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::DictionaryField<std::string>>& item);
bool VectorDictionaryFieldStringWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryField<std::string>>& item);

}} // namespace tl2::details

