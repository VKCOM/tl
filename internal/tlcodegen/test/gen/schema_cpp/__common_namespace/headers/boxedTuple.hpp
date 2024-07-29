#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/boxedTuple.hpp"
#include "../types/int.hpp"

namespace tl2 { namespace details { 

void BoxedTupleReset(::tl2::BoxedTuple& item);
bool BoxedTupleRead(::basictl::tl_istream & s, ::tl2::BoxedTuple& item);
bool BoxedTupleWrite(::basictl::tl_ostream & s, const ::tl2::BoxedTuple& item);
bool BoxedTupleReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedTuple& item);
bool BoxedTupleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedTuple& item);

bool BoxedTupleReadResult(::basictl::tl_istream & s, ::tl2::BoxedTuple& item, std::array<int32_t, 3>& result);
bool BoxedTupleWriteResult(::basictl::tl_ostream & s, ::tl2::BoxedTuple& item, std::array<int32_t, 3>& result);
		
}} // namespace tl2::details

