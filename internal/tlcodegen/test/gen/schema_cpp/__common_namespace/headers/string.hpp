#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/string.hpp"

namespace tl2 { namespace details { 

void BuiltinVectorStringReset(std::vector<std::string>& item);

bool BuiltinVectorStringWriteJSON(std::ostream & s, const std::vector<std::string>& item);
bool BuiltinVectorStringRead(::basictl::tl_istream & s, std::vector<std::string>& item);
bool BuiltinVectorStringWrite(::basictl::tl_ostream & s, const std::vector<std::string>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void StringReset(std::string& item);

bool StringWriteJSON(std::ostream& s, const std::string& item);
bool StringRead(::basictl::tl_istream & s, std::string& item);
bool StringWrite(::basictl::tl_ostream & s, const std::string& item);
bool StringReadBoxed(::basictl::tl_istream & s, std::string& item);
bool StringWriteBoxed(::basictl::tl_ostream & s, const std::string& item);

}} // namespace tl2::details

