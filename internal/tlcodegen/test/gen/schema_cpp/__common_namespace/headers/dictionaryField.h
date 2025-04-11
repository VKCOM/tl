#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/dictionaryField.h"
#include "../../service1/types/service1.Value.h"

namespace tl2 { namespace details { 

void BuiltinVectorDictionaryFieldIntReset(std::map<std::string, int32_t>& item);

bool BuiltinVectorDictionaryFieldIntWriteJSON(std::ostream & s, const std::map<std::string, int32_t>& item);
bool BuiltinVectorDictionaryFieldIntRead(::basictl::tl_istream & s, std::map<std::string, int32_t>& item);
bool BuiltinVectorDictionaryFieldIntWrite(::basictl::tl_ostream & s, const std::map<std::string, int32_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorDictionaryFieldService1ValueReset(std::map<std::string, ::tl2::service1::Value>& item);

bool BuiltinVectorDictionaryFieldService1ValueWriteJSON(std::ostream & s, const std::map<std::string, ::tl2::service1::Value>& item);
bool BuiltinVectorDictionaryFieldService1ValueRead(::basictl::tl_istream & s, std::map<std::string, ::tl2::service1::Value>& item);
bool BuiltinVectorDictionaryFieldService1ValueWrite(::basictl::tl_ostream & s, const std::map<std::string, ::tl2::service1::Value>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorDictionaryFieldStringReset(std::map<std::string, std::string>& item);

bool BuiltinVectorDictionaryFieldStringWriteJSON(std::ostream & s, const std::map<std::string, std::string>& item);
bool BuiltinVectorDictionaryFieldStringRead(::basictl::tl_istream & s, std::map<std::string, std::string>& item);
bool BuiltinVectorDictionaryFieldStringWrite(::basictl::tl_ostream & s, const std::map<std::string, std::string>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryFieldDictionaryIntReset(::tl2::DictionaryField<std::map<std::string, int32_t>>& item) noexcept;

bool DictionaryFieldDictionaryIntWriteJSON(std::ostream& s, const ::tl2::DictionaryField<std::map<std::string, int32_t>>& item) noexcept;
bool DictionaryFieldDictionaryIntRead(::basictl::tl_istream & s, ::tl2::DictionaryField<std::map<std::string, int32_t>>& item) noexcept; 
bool DictionaryFieldDictionaryIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryField<std::map<std::string, int32_t>>& item) noexcept;
bool DictionaryFieldDictionaryIntReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryField<std::map<std::string, int32_t>>& item);
bool DictionaryFieldDictionaryIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryField<std::map<std::string, int32_t>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryFieldIntReset(::tl2::DictionaryField<int32_t>& item) noexcept;

bool DictionaryFieldIntWriteJSON(std::ostream& s, const ::tl2::DictionaryField<int32_t>& item) noexcept;
bool DictionaryFieldIntRead(::basictl::tl_istream & s, ::tl2::DictionaryField<int32_t>& item) noexcept; 
bool DictionaryFieldIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryField<int32_t>& item) noexcept;
bool DictionaryFieldIntReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryField<int32_t>& item);
bool DictionaryFieldIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryField<int32_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryFieldService1ValueReset(::tl2::DictionaryField<::tl2::service1::Value>& item) noexcept;

bool DictionaryFieldService1ValueWriteJSON(std::ostream& s, const ::tl2::DictionaryField<::tl2::service1::Value>& item) noexcept;
bool DictionaryFieldService1ValueRead(::basictl::tl_istream & s, ::tl2::DictionaryField<::tl2::service1::Value>& item) noexcept; 
bool DictionaryFieldService1ValueWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryField<::tl2::service1::Value>& item) noexcept;
bool DictionaryFieldService1ValueReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryField<::tl2::service1::Value>& item);
bool DictionaryFieldService1ValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryField<::tl2::service1::Value>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryFieldStringReset(::tl2::DictionaryField<std::string>& item) noexcept;

bool DictionaryFieldStringWriteJSON(std::ostream& s, const ::tl2::DictionaryField<std::string>& item) noexcept;
bool DictionaryFieldStringRead(::basictl::tl_istream & s, ::tl2::DictionaryField<std::string>& item) noexcept; 
bool DictionaryFieldStringWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryField<std::string>& item) noexcept;
bool DictionaryFieldStringReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryField<std::string>& item);
bool DictionaryFieldStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryField<std::string>& item);

}} // namespace tl2::details

