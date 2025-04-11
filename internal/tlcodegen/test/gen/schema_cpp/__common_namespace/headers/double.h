#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/double.h"

namespace tl2 { namespace details { 

void DoubleReset(double& item) noexcept;

bool DoubleWriteJSON(std::ostream& s, const double& item) noexcept;
bool DoubleRead(::basictl::tl_istream & s, double& item) noexcept; 
bool DoubleWrite(::basictl::tl_ostream & s, const double& item) noexcept;
bool DoubleReadBoxed(::basictl::tl_istream & s, double& item);
bool DoubleWriteBoxed(::basictl::tl_ostream & s, const double& item);

}} // namespace tl2::details

