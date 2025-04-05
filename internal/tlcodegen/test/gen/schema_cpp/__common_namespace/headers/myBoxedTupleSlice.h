#pragma once

#include "../../basictl/io_streams.h"
#include "../types/myBoxedTupleSlice.h"

namespace tl2 { namespace details { 

void MyBoxedTupleSliceReset(::tl2::MyBoxedTupleSlice& item);

bool MyBoxedTupleSliceWriteJSON(std::ostream& s, const ::tl2::MyBoxedTupleSlice& item);
bool MyBoxedTupleSliceRead(::basictl::tl_istream & s, ::tl2::MyBoxedTupleSlice& item);
bool MyBoxedTupleSliceWrite(::basictl::tl_ostream & s, const ::tl2::MyBoxedTupleSlice& item);
bool MyBoxedTupleSliceReadBoxed(::basictl::tl_istream & s, ::tl2::MyBoxedTupleSlice& item);
bool MyBoxedTupleSliceWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyBoxedTupleSlice& item);

}} // namespace tl2::details

