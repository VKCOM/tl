// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/vector.h"
#include "__common_namespace/types/string.h"
#include "__common_namespace/types/dictionaryField.h"

namespace tlgen { namespace details { 

void VectorDictionaryFieldStringReset(std::map<std::string, std::string>& item) noexcept;

bool VectorDictionaryFieldStringWriteJSON(std::ostream& s, const std::map<std::string, std::string>& item) noexcept;
bool VectorDictionaryFieldStringRead(::tlgen::basictl::tl_istream & s, std::map<std::string, std::string>& item) noexcept; 
bool VectorDictionaryFieldStringWrite(::tlgen::basictl::tl_ostream & s, const std::map<std::string, std::string>& item) noexcept;
bool VectorDictionaryFieldStringReadBoxed(::tlgen::basictl::tl_istream & s, std::map<std::string, std::string>& item);
bool VectorDictionaryFieldStringWriteBoxed(::tlgen::basictl::tl_ostream & s, const std::map<std::string, std::string>& item);

}} // namespace tlgen::details

namespace tlgen { namespace details { 

void VectorStringReset(std::vector<std::string>& item) noexcept;

bool VectorStringWriteJSON(std::ostream& s, const std::vector<std::string>& item) noexcept;
bool VectorStringRead(::tlgen::basictl::tl_istream & s, std::vector<std::string>& item) noexcept; 
bool VectorStringWrite(::tlgen::basictl::tl_ostream & s, const std::vector<std::string>& item) noexcept;
bool VectorStringReadBoxed(::tlgen::basictl::tl_istream & s, std::vector<std::string>& item);
bool VectorStringWriteBoxed(::tlgen::basictl::tl_ostream & s, const std::vector<std::string>& item);

}} // namespace tlgen::details

