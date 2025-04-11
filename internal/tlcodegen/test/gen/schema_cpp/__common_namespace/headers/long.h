#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/long.h"

namespace tl2 { namespace details { 

void BuiltinVectorLongBoxedReset(std::vector<int64_t>& item);

bool BuiltinVectorLongBoxedWriteJSON(std::ostream & s, const std::vector<int64_t>& item);
bool BuiltinVectorLongBoxedRead(::basictl::tl_istream & s, std::vector<int64_t>& item);
bool BuiltinVectorLongBoxedWrite(::basictl::tl_ostream & s, const std::vector<int64_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void LongReset(int64_t& item) noexcept;

bool LongWriteJSON(std::ostream& s, const int64_t& item) noexcept;
bool LongRead(::basictl::tl_istream & s, int64_t& item) noexcept; 
bool LongWrite(::basictl::tl_ostream & s, const int64_t& item) noexcept;
bool LongReadBoxed(::basictl::tl_istream & s, int64_t& item);
bool LongWriteBoxed(::basictl::tl_ostream & s, const int64_t& item);

}} // namespace tl2::details

