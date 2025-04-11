#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/service6.error.h"

namespace tl2 { namespace details { 

void Service6ErrorReset(::tl2::service6::Error& item) noexcept;

bool Service6ErrorWriteJSON(std::ostream& s, const ::tl2::service6::Error& item) noexcept;
bool Service6ErrorRead(::basictl::tl_istream & s, ::tl2::service6::Error& item) noexcept; 
bool Service6ErrorWrite(::basictl::tl_ostream & s, const ::tl2::service6::Error& item) noexcept;
bool Service6ErrorReadBoxed(::basictl::tl_istream & s, ::tl2::service6::Error& item);
bool Service6ErrorWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service6::Error& item);

}} // namespace tl2::details

