#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../__common_namespace/types/tuple.hpp"
#include "../../__common_namespace/types/int.hpp"

namespace tl2 { namespace details { 

void TupleInt4Reset(std::array<int32_t, 4>& item);

bool TupleInt4WriteJSON(std::ostream& s, const std::array<int32_t, 4>& item);
bool TupleInt4Read(::basictl::tl_istream & s, std::array<int32_t, 4>& item);
bool TupleInt4Write(::basictl::tl_ostream & s, const std::array<int32_t, 4>& item);
bool TupleInt4ReadBoxed(::basictl::tl_istream & s, std::array<int32_t, 4>& item);
bool TupleInt4WriteBoxed(::basictl::tl_ostream & s, const std::array<int32_t, 4>& item);

}} // namespace tl2::details

