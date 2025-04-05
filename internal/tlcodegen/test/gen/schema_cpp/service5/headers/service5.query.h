#pragma once

#include "../../basictl/io_streams.h"
#include "../functions/service5.query.h"
#include "../types/service5.Output.h"

namespace tl2 { namespace details { 

void Service5QueryReset(::tl2::service5::Query& item);

bool Service5QueryWriteJSON(std::ostream& s, const ::tl2::service5::Query& item);
bool Service5QueryRead(::basictl::tl_istream & s, ::tl2::service5::Query& item);
bool Service5QueryWrite(::basictl::tl_ostream & s, const ::tl2::service5::Query& item);
bool Service5QueryReadBoxed(::basictl::tl_istream & s, ::tl2::service5::Query& item);
bool Service5QueryWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service5::Query& item);

bool Service5QueryReadResult(::basictl::tl_istream & s, ::tl2::service5::Query& item, ::tl2::service5::Output& result);
bool Service5QueryWriteResult(::basictl::tl_ostream & s, ::tl2::service5::Query& item, ::tl2::service5::Output& result);
		
}} // namespace tl2::details

