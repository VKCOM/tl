#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/string.h"

namespace tl2 { namespace details { 

void StringReset(std::string& item) noexcept;

bool StringWriteJSON(std::ostream& s, const std::string& item) noexcept;
bool StringRead(::basictl::tl_istream & s, std::string& item) noexcept; 
bool StringWrite(::basictl::tl_ostream & s, const std::string& item) noexcept;
bool StringReadBoxed(::basictl::tl_istream & s, std::string& item);
bool StringWriteBoxed(::basictl::tl_ostream & s, const std::string& item);

}} // namespace tl2::details

