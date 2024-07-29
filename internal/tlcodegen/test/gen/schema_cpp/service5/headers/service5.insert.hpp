#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/service5.insert.hpp"
#include "../types/service5.Output.hpp"

namespace tl2 { namespace details { 

void Service5InsertReset(::tl2::service5::Insert& item);
bool Service5InsertRead(::basictl::tl_istream & s, ::tl2::service5::Insert& item);
bool Service5InsertWrite(::basictl::tl_ostream & s, const ::tl2::service5::Insert& item);
bool Service5InsertReadBoxed(::basictl::tl_istream & s, ::tl2::service5::Insert& item);
bool Service5InsertWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service5::Insert& item);

bool Service5InsertReadResult(::basictl::tl_istream & s, ::tl2::service5::Insert& item, ::tl2::service5::Output& result);
bool Service5InsertWriteResult(::basictl::tl_ostream & s, ::tl2::service5::Insert& item, ::tl2::service5::Output& result);
		
}} // namespace tl2::details

