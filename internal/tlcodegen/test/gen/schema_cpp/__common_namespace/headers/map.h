#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/map.h"

namespace tl2 { namespace details { 

void BuiltinVectorMapStringStringReset(std::vector<::tl2::Map<std::string, std::string>>& item);

bool BuiltinVectorMapStringStringWriteJSON(std::ostream & s, const std::vector<::tl2::Map<std::string, std::string>>& item);
bool BuiltinVectorMapStringStringRead(::basictl::tl_istream & s, std::vector<::tl2::Map<std::string, std::string>>& item);
bool BuiltinVectorMapStringStringWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Map<std::string, std::string>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void MapStringStringReset(::tl2::Map<std::string, std::string>& item);

bool MapStringStringWriteJSON(std::ostream& s, const ::tl2::Map<std::string, std::string>& item);
bool MapStringStringRead(::basictl::tl_istream & s, ::tl2::Map<std::string, std::string>& item);
bool MapStringStringWrite(::basictl::tl_ostream & s, const ::tl2::Map<std::string, std::string>& item);
bool MapStringStringReadBoxed(::basictl::tl_istream & s, ::tl2::Map<std::string, std::string>& item);
bool MapStringStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Map<std::string, std::string>& item);

}} // namespace tl2::details

