#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service5.emptyOutput.hpp"

namespace tl2 { namespace details { 

void Service5EmptyOutputReset(::tl2::service5::EmptyOutput& item);
bool Service5EmptyOutputRead(::basictl::tl_istream & s, ::tl2::service5::EmptyOutput& item);
bool Service5EmptyOutputWrite(::basictl::tl_ostream & s, const ::tl2::service5::EmptyOutput& item);
bool Service5EmptyOutputReadBoxed(::basictl::tl_istream & s, ::tl2::service5::EmptyOutput& item);
bool Service5EmptyOutputWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service5::EmptyOutput& item);

}} // namespace tl2::details

