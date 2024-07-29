#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/service5.performQuery.hpp"
#include "../types/service5.Output.hpp"

namespace tl2 { namespace details { 

void Service5PerformQueryReset(::tl2::service5::PerformQuery& item);
bool Service5PerformQueryRead(::basictl::tl_istream & s, ::tl2::service5::PerformQuery& item);
bool Service5PerformQueryWrite(::basictl::tl_ostream & s, const ::tl2::service5::PerformQuery& item);
bool Service5PerformQueryReadBoxed(::basictl::tl_istream & s, ::tl2::service5::PerformQuery& item);
bool Service5PerformQueryWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service5::PerformQuery& item);

bool Service5PerformQueryReadResult(::basictl::tl_istream & s, ::tl2::service5::PerformQuery& item, ::tl2::service5::Output& result);
bool Service5PerformQueryWriteResult(::basictl::tl_ostream & s, ::tl2::service5::PerformQuery& item, ::tl2::service5::Output& result);
		
}} // namespace tl2::details

