#pragma once

#include "../../basictl/io_streams.h"
#include "../functions/boxedTuple.h"
#include "../types/int.h"

namespace tl2 { namespace details { 

void BoxedTupleReset(::tl2::BoxedTuple& item);

bool BoxedTupleWriteJSON(std::ostream& s, const ::tl2::BoxedTuple& item);
bool BoxedTupleRead(::basictl::tl_istream & s, ::tl2::BoxedTuple& item);
bool BoxedTupleWrite(::basictl::tl_ostream & s, const ::tl2::BoxedTuple& item);
bool BoxedTupleReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedTuple& item);
bool BoxedTupleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedTuple& item);

bool BoxedTupleReadResult(::basictl::tl_istream & s, ::tl2::BoxedTuple& item, std::array<int32_t, 3>& result);
bool BoxedTupleWriteResult(::basictl::tl_ostream & s, ::tl2::BoxedTuple& item, std::array<int32_t, 3>& result);
		
}} // namespace tl2::details

