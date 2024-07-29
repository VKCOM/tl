#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../../__common/types/dictionary.hpp"
#include "../../types/service1.Value.hpp"

namespace tl2 { namespace details { 

void DictionaryDictionaryIntReset(::tl2::Dictionary<::tl2::Dictionary<int32_t>>& item);
bool DictionaryDictionaryIntRead(::basictl::tl_istream & s, ::tl2::Dictionary<::tl2::Dictionary<int32_t>>& item);
bool DictionaryDictionaryIntWrite(::basictl::tl_ostream & s, const ::tl2::Dictionary<::tl2::Dictionary<int32_t>>& item);
bool DictionaryDictionaryIntReadBoxed(::basictl::tl_istream & s, ::tl2::Dictionary<::tl2::Dictionary<int32_t>>& item);
bool DictionaryDictionaryIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Dictionary<::tl2::Dictionary<int32_t>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryService1ValueReset(::tl2::Dictionary<::tl2::service1::Value>& item);
bool DictionaryService1ValueRead(::basictl::tl_istream & s, ::tl2::Dictionary<::tl2::service1::Value>& item);
bool DictionaryService1ValueWrite(::basictl::tl_ostream & s, const ::tl2::Dictionary<::tl2::service1::Value>& item);
bool DictionaryService1ValueReadBoxed(::basictl::tl_istream & s, ::tl2::Dictionary<::tl2::service1::Value>& item);
bool DictionaryService1ValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Dictionary<::tl2::service1::Value>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryStringReset(::tl2::Dictionary<std::string>& item);
bool DictionaryStringRead(::basictl::tl_istream & s, ::tl2::Dictionary<std::string>& item);
bool DictionaryStringWrite(::basictl::tl_ostream & s, const ::tl2::Dictionary<std::string>& item);
bool DictionaryStringReadBoxed(::basictl::tl_istream & s, ::tl2::Dictionary<std::string>& item);
bool DictionaryStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Dictionary<std::string>& item);

}} // namespace tl2::details

