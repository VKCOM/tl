#pragma once

#include "../../basictl/io_streams.h"
#include "../types/service5.params.h"

namespace tl2 { namespace details { 

void Service5ParamsReset(::tl2::service5::Params& item);

bool Service5ParamsWriteJSON(std::ostream& s, const ::tl2::service5::Params& item);
bool Service5ParamsRead(::basictl::tl_istream & s, ::tl2::service5::Params& item);
bool Service5ParamsWrite(::basictl::tl_ostream & s, const ::tl2::service5::Params& item);
bool Service5ParamsReadBoxed(::basictl::tl_istream & s, ::tl2::service5::Params& item);
bool Service5ParamsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service5::Params& item);

}} // namespace tl2::details

