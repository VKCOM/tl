#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service5.Output.hpp"

namespace tl2 { namespace details { 

void Service5OutputReset(::tl2::service5::Output& item);
bool Service5OutputReadBoxed(::basictl::tl_istream & s, ::tl2::service5::Output& item);
bool Service5OutputWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service5::Output& item);

}} // namespace tl2::details

