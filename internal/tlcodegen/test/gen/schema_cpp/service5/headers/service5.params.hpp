#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service5.params.hpp"

namespace tl2 { namespace details { 

void Service5ParamsReset(::tl2::service5::Params& item);

bool Service5ParamsWriteJSON(std::ostream& s, const ::tl2::service5::Params& item);
bool Service5ParamsRead(::basictl::tl_istream & s, ::tl2::service5::Params& item);
bool Service5ParamsWrite(::basictl::tl_ostream & s, const ::tl2::service5::Params& item);
bool Service5ParamsReadBoxed(::basictl::tl_istream & s, ::tl2::service5::Params& item);
bool Service5ParamsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service5::Params& item);

}} // namespace tl2::details

