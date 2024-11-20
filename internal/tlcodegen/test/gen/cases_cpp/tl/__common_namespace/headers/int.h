#pragma once

#include "../../../basics/basictl.h"
#include "../types/int.h"

namespace tl2 { namespace details { 

void IntReset(int32_t& item);

bool IntWriteJSON(std::ostream& s, const int32_t& item);
bool IntRead(::basictl::tl_istream & s, int32_t& item);
bool IntWrite(::basictl::tl_ostream & s, const int32_t& item);
bool IntReadBoxed(::basictl::tl_istream & s, int32_t& item);
bool IntWriteBoxed(::basictl::tl_ostream & s, const int32_t& item);

}} // namespace tl2::details

