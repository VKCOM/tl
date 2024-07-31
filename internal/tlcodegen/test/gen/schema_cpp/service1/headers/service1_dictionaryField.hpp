#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../__common_namespace/types/dictionaryField.hpp"
#include "../types/service1.Value.hpp"
#include "../../__common_namespace/types/dictionary.hpp"

namespace tl2 { namespace details { 

void BuiltinVectorDictionaryFieldDictionaryIntReset(std::vector<::tl2::DictionaryField<::tl2::Dictionary<int32_t>>>& item);

bool BuiltinVectorDictionaryFieldDictionaryIntWriteJSON(std::ostream & s, const std::vector<::tl2::DictionaryField<::tl2::Dictionary<int32_t>>>& item);
bool BuiltinVectorDictionaryFieldDictionaryIntRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryField<::tl2::Dictionary<int32_t>>>& item);
bool BuiltinVectorDictionaryFieldDictionaryIntWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryField<::tl2::Dictionary<int32_t>>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorDictionaryFieldService1ValueReset(std::vector<::tl2::DictionaryField<::tl2::service1::Value>>& item);

bool BuiltinVectorDictionaryFieldService1ValueWriteJSON(std::ostream & s, const std::vector<::tl2::DictionaryField<::tl2::service1::Value>>& item);
bool BuiltinVectorDictionaryFieldService1ValueRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryField<::tl2::service1::Value>>& item);
bool BuiltinVectorDictionaryFieldService1ValueWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryField<::tl2::service1::Value>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorDictionaryFieldStringReset(std::vector<::tl2::DictionaryField<std::string>>& item);

bool BuiltinVectorDictionaryFieldStringWriteJSON(std::ostream & s, const std::vector<::tl2::DictionaryField<std::string>>& item);
bool BuiltinVectorDictionaryFieldStringRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryField<std::string>>& item);
bool BuiltinVectorDictionaryFieldStringWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryField<std::string>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryFieldDictionaryIntReset(::tl2::DictionaryField<::tl2::Dictionary<int32_t>>& item);

bool DictionaryFieldDictionaryIntWriteJSON(std::ostream& s, const ::tl2::DictionaryField<::tl2::Dictionary<int32_t>>& item);
bool DictionaryFieldDictionaryIntRead(::basictl::tl_istream & s, ::tl2::DictionaryField<::tl2::Dictionary<int32_t>>& item);
bool DictionaryFieldDictionaryIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryField<::tl2::Dictionary<int32_t>>& item);
bool DictionaryFieldDictionaryIntReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryField<::tl2::Dictionary<int32_t>>& item);
bool DictionaryFieldDictionaryIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryField<::tl2::Dictionary<int32_t>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryFieldService1ValueReset(::tl2::DictionaryField<::tl2::service1::Value>& item);

bool DictionaryFieldService1ValueWriteJSON(std::ostream& s, const ::tl2::DictionaryField<::tl2::service1::Value>& item);
bool DictionaryFieldService1ValueRead(::basictl::tl_istream & s, ::tl2::DictionaryField<::tl2::service1::Value>& item);
bool DictionaryFieldService1ValueWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryField<::tl2::service1::Value>& item);
bool DictionaryFieldService1ValueReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryField<::tl2::service1::Value>& item);
bool DictionaryFieldService1ValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryField<::tl2::service1::Value>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryFieldStringReset(::tl2::DictionaryField<std::string>& item);

bool DictionaryFieldStringWriteJSON(std::ostream& s, const ::tl2::DictionaryField<std::string>& item);
bool DictionaryFieldStringRead(::basictl::tl_istream & s, ::tl2::DictionaryField<std::string>& item);
bool DictionaryFieldStringWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryField<std::string>& item);
bool DictionaryFieldStringReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryField<std::string>& item);
bool DictionaryFieldStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryField<std::string>& item);

}} // namespace tl2::details

