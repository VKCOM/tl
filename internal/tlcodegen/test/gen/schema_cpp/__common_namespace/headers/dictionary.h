#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/dictionary.h"
#include "../../service1/types/service1.Value.h"
#include "../types/dictionaryField.h"

namespace tl2 { namespace details { 

void DictionaryIntReset(std::map<std::string, int32_t>& item) noexcept;

bool DictionaryIntWriteJSON(std::ostream& s, const std::map<std::string, int32_t>& item) noexcept;
bool DictionaryIntRead(::basictl::tl_istream & s, std::map<std::string, int32_t>& item) noexcept; 
bool DictionaryIntWrite(::basictl::tl_ostream & s, const std::map<std::string, int32_t>& item) noexcept;
bool DictionaryIntReadBoxed(::basictl::tl_istream & s, std::map<std::string, int32_t>& item);
bool DictionaryIntWriteBoxed(::basictl::tl_ostream & s, const std::map<std::string, int32_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryService1ValueReset(std::map<std::string, ::tl2::service1::Value>& item) noexcept;

bool DictionaryService1ValueWriteJSON(std::ostream& s, const std::map<std::string, ::tl2::service1::Value>& item) noexcept;
bool DictionaryService1ValueRead(::basictl::tl_istream & s, std::map<std::string, ::tl2::service1::Value>& item) noexcept; 
bool DictionaryService1ValueWrite(::basictl::tl_ostream & s, const std::map<std::string, ::tl2::service1::Value>& item) noexcept;
bool DictionaryService1ValueReadBoxed(::basictl::tl_istream & s, std::map<std::string, ::tl2::service1::Value>& item);
bool DictionaryService1ValueWriteBoxed(::basictl::tl_ostream & s, const std::map<std::string, ::tl2::service1::Value>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryStringReset(std::map<std::string, std::string>& item) noexcept;

bool DictionaryStringWriteJSON(std::ostream& s, const std::map<std::string, std::string>& item) noexcept;
bool DictionaryStringRead(::basictl::tl_istream & s, std::map<std::string, std::string>& item) noexcept; 
bool DictionaryStringWrite(::basictl::tl_ostream & s, const std::map<std::string, std::string>& item) noexcept;
bool DictionaryStringReadBoxed(::basictl::tl_istream & s, std::map<std::string, std::string>& item);
bool DictionaryStringWriteBoxed(::basictl::tl_ostream & s, const std::map<std::string, std::string>& item);

}} // namespace tl2::details

