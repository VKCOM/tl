#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/int.h"

namespace tl2 { namespace details { 

void IntReset(int32_t& item) noexcept;

bool IntWriteJSON(std::ostream& s, const int32_t& item) noexcept;
bool IntRead(::basictl::tl_istream & s, int32_t& item) noexcept; 
bool IntWrite(::basictl::tl_ostream & s, const int32_t& item) noexcept;
bool IntReadBoxed(::basictl::tl_istream & s, int32_t& item);
bool IntWriteBoxed(::basictl::tl_ostream & s, const int32_t& item);

}} // namespace tl2::details

