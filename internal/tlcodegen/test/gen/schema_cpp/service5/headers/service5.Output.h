#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/service5.Output.h"

namespace tl2 { namespace details { 

void Service5OutputReset(::tl2::service5::Output& item) noexcept;

bool Service5OutputWriteJSON(std::ostream & s, const ::tl2::service5::Output& item) noexcept;
bool Service5OutputReadBoxed(::basictl::tl_istream & s, ::tl2::service5::Output& item) noexcept;
bool Service5OutputWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service5::Output& item) noexcept;

}} // namespace tl2::details

