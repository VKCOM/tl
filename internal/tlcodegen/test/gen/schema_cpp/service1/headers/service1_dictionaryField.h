#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/dictionaryField.h"
#include "service1/types/service1.Value.h"

namespace tl2 { namespace details { 

void BuiltinVectorDictionaryFieldDictionaryIntReset(std::map<std::string, std::map<std::string, int32_t>>& item);

bool BuiltinVectorDictionaryFieldDictionaryIntWriteJSON(std::ostream & s, const std::map<std::string, std::map<std::string, int32_t>>& item);
bool BuiltinVectorDictionaryFieldDictionaryIntRead(::basictl::tl_istream & s, std::map<std::string, std::map<std::string, int32_t>>& item);
bool BuiltinVectorDictionaryFieldDictionaryIntWrite(::basictl::tl_ostream & s, const std::map<std::string, std::map<std::string, int32_t>>& item);

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

