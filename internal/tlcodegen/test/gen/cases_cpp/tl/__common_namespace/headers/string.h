#pragma once

#include "../../../basics/basictl.h"
#include "../types/string.h"

namespace tl2 { namespace details { 

void StringReset(std::string& item);

bool StringWriteJSON(std::ostream& s, const std::string& item);
bool StringRead(::basictl::tl_istream & s, std::string& item);
bool StringWrite(::basictl::tl_ostream & s, const std::string& item);
bool StringReadBoxed(::basictl::tl_istream & s, std::string& item);
bool StringWriteBoxed(::basictl::tl_ostream & s, const std::string& item);

}} // namespace tl2::details

