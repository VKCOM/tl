#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/MyValue.h"

namespace tl2 { namespace details { 

void MyValueReset(::tl2::MyValue& item) noexcept;

bool MyValueWriteJSON(std::ostream & s, const ::tl2::MyValue& item) noexcept;
bool MyValueReadBoxed(::basictl::tl_istream & s, ::tl2::MyValue& item) noexcept;
bool MyValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyValue& item) noexcept;

}} // namespace tl2::details

