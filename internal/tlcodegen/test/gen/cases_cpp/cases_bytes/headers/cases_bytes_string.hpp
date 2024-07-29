#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../__common_namespace/types/string.hpp"

namespace tl2 { namespace details { 

void BuiltinTuple4StringReset(std::array<std::string, 4>& item);
bool BuiltinTuple4StringRead(::basictl::tl_istream & s, std::array<std::string, 4>& item);
bool BuiltinTuple4StringWrite(::basictl::tl_ostream & s, const std::array<std::string, 4>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinTupleStringReset(std::vector<std::string>& item);
bool BuiltinTupleStringRead(::basictl::tl_istream & s, std::vector<std::string>& item, uint32_t nat_n);
bool BuiltinTupleStringWrite(::basictl::tl_ostream & s, const std::vector<std::string>& item, uint32_t nat_n);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorStringReset(std::vector<std::string>& item);
bool BuiltinVectorStringRead(::basictl::tl_istream & s, std::vector<std::string>& item);
bool BuiltinVectorStringWrite(::basictl::tl_ostream & s, const std::vector<std::string>& item);

}} // namespace tl2::details

