#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/double.hpp"

namespace tl2 { namespace details { 

void DoubleReset(double& item);

bool DoubleWriteJSON(std::ostream& s, const double& item);
bool DoubleRead(::basictl::tl_istream & s, double& item);
bool DoubleWrite(::basictl::tl_ostream & s, const double& item);
bool DoubleReadBoxed(::basictl::tl_istream & s, double& item);
bool DoubleWriteBoxed(::basictl::tl_ostream & s, const double& item);

}} // namespace tl2::details

