#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/long.hpp"

namespace tl2 { namespace details { 

void LongReset(int64_t& item);

bool LongWriteJSON(std::ostream& s, const int64_t& item);
bool LongRead(::basictl::tl_istream & s, int64_t& item);
bool LongWrite(::basictl::tl_ostream & s, const int64_t& item);
bool LongReadBoxed(::basictl::tl_istream & s, int64_t& item);
bool LongWriteBoxed(::basictl::tl_ostream & s, const int64_t& item);

}} // namespace tl2::details

