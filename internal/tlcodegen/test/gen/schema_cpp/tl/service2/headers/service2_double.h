#pragma once

#include "../../../basics/basictl.h"
#include "../../__common_namespace/types/double.h"

namespace tl2 { namespace details { 

void BuiltinTupleDoubleReset(std::vector<double>& item);

bool BuiltinTupleDoubleWriteJSON(std::ostream & s, const std::vector<double>& item, uint32_t nat_n);
bool BuiltinTupleDoubleRead(::basictl::tl_istream & s, std::vector<double>& item, uint32_t nat_n);
bool BuiltinTupleDoubleWrite(::basictl::tl_ostream & s, const std::vector<double>& item, uint32_t nat_n);

}} // namespace tl2::details

