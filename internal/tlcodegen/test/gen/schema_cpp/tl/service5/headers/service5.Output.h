#pragma once

#include "../../../basics/basictl.h"
#include "../types/service5.Output.h"

namespace tl2 { namespace details { 

void Service5OutputReset(::tl2::service5::Output& item);

bool Service5OutputWriteJSON(std::ostream & s, const ::tl2::service5::Output& item);
bool Service5OutputReadBoxed(::basictl::tl_istream & s, ::tl2::service5::Output& item);
bool Service5OutputWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service5::Output& item);

}} // namespace tl2::details

