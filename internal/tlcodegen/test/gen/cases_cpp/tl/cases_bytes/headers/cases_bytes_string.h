#pragma once

#include "../../../basics/basictl.h"
#include "../../__common_namespace/types/string.h"

namespace tl2 { namespace details { 

void BuiltinTuple4StringReset(std::array<std::string, 4>& item);

bool BuiltinTuple4StringWriteJSON(std::ostream & s, const std::array<std::string, 4>& item);
bool BuiltinTuple4StringRead(::basictl::tl_istream & s, std::array<std::string, 4>& item);
bool BuiltinTuple4StringWrite(::basictl::tl_ostream & s, const std::array<std::string, 4>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinTupleStringReset(std::vector<std::string>& item);

bool BuiltinTupleStringWriteJSON(std::ostream & s, const std::vector<std::string>& item, uint32_t nat_n);
bool BuiltinTupleStringRead(::basictl::tl_istream & s, std::vector<std::string>& item, uint32_t nat_n);
bool BuiltinTupleStringWrite(::basictl::tl_ostream & s, const std::vector<std::string>& item, uint32_t nat_n);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorStringReset(std::vector<std::string>& item);

bool BuiltinVectorStringWriteJSON(std::ostream & s, const std::vector<std::string>& item);
bool BuiltinVectorStringRead(::basictl::tl_istream & s, std::vector<std::string>& item);
bool BuiltinVectorStringWrite(::basictl::tl_ostream & s, const std::vector<std::string>& item);

}} // namespace tl2::details

