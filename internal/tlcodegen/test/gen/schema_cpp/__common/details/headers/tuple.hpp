#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../types/tuple.hpp"
#include "../../types/int.hpp"

namespace tl2 { namespace details { 

void TupleIntReset(std::vector<int32_t>& item);
bool TupleIntRead(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n);
bool TupleIntWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n);
bool TupleIntReadBoxed(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n);
bool TupleIntWriteBoxed(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n);

}} // namespace tl2::details

