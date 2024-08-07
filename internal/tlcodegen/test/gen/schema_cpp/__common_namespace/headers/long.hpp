#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/long.hpp"

namespace tl2 { namespace details { 

void BuiltinVectorLongBoxedReset(std::vector<int64_t>& item);

bool BuiltinVectorLongBoxedWriteJSON(std::ostream & s, const std::vector<int64_t>& item);
bool BuiltinVectorLongBoxedRead(::basictl::tl_istream & s, std::vector<int64_t>& item);
bool BuiltinVectorLongBoxedWrite(::basictl::tl_ostream & s, const std::vector<int64_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void LongReset(int64_t& item);

bool LongWriteJSON(std::ostream& s, const int64_t& item);
bool LongRead(::basictl::tl_istream & s, int64_t& item);
bool LongWrite(::basictl::tl_ostream & s, const int64_t& item);
bool LongReadBoxed(::basictl::tl_istream & s, int64_t& item);
bool LongWriteBoxed(::basictl::tl_ostream & s, const int64_t& item);

}} // namespace tl2::details

