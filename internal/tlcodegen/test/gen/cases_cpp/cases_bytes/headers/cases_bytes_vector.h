// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/vector.h"
#include "__common_namespace/types/string.h"
#include "__common_namespace/types/dictionaryField.h"

namespace tl2 { namespace details { 

void VectorDictionaryFieldStringReset(std::map<std::string, std::string>& item) noexcept;

bool VectorDictionaryFieldStringWriteJSON(std::ostream& s, const std::map<std::string, std::string>& item) noexcept;
bool VectorDictionaryFieldStringRead(::basictl::tl_istream & s, std::map<std::string, std::string>& item) noexcept; 
bool VectorDictionaryFieldStringWrite(::basictl::tl_ostream & s, const std::map<std::string, std::string>& item) noexcept;
bool VectorDictionaryFieldStringReadBoxed(::basictl::tl_istream & s, std::map<std::string, std::string>& item);
bool VectorDictionaryFieldStringWriteBoxed(::basictl::tl_ostream & s, const std::map<std::string, std::string>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorStringReset(std::vector<std::string>& item) noexcept;

bool VectorStringWriteJSON(std::ostream& s, const std::vector<std::string>& item) noexcept;
bool VectorStringRead(::basictl::tl_istream & s, std::vector<std::string>& item) noexcept; 
bool VectorStringWrite(::basictl::tl_ostream & s, const std::vector<std::string>& item) noexcept;
bool VectorStringReadBoxed(::basictl::tl_istream & s, std::vector<std::string>& item);
bool VectorStringWriteBoxed(::basictl::tl_ostream & s, const std::vector<std::string>& item);

}} // namespace tl2::details

