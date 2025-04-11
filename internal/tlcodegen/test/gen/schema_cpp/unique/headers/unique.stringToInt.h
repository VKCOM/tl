#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/unique.stringToInt.h"

namespace tl2 { namespace details { 

void UniqueStringToIntReset(::tl2::unique::StringToInt& item) noexcept;

bool UniqueStringToIntWriteJSON(std::ostream& s, const ::tl2::unique::StringToInt& item) noexcept;
bool UniqueStringToIntRead(::basictl::tl_istream & s, ::tl2::unique::StringToInt& item) noexcept; 
bool UniqueStringToIntWrite(::basictl::tl_ostream & s, const ::tl2::unique::StringToInt& item) noexcept;
bool UniqueStringToIntReadBoxed(::basictl::tl_istream & s, ::tl2::unique::StringToInt& item);
bool UniqueStringToIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::unique::StringToInt& item);

bool UniqueStringToIntReadResult(::basictl::tl_istream & s, ::tl2::unique::StringToInt& item, int32_t& result);
bool UniqueStringToIntWriteResult(::basictl::tl_ostream & s, ::tl2::unique::StringToInt& item, int32_t& result);
		
}} // namespace tl2::details

