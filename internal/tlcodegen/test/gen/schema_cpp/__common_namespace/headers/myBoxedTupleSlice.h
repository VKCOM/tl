#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/myBoxedTupleSlice.h"

namespace tl2 { namespace details { 

void MyBoxedTupleSliceReset(::tl2::MyBoxedTupleSlice& item) noexcept;

bool MyBoxedTupleSliceWriteJSON(std::ostream& s, const ::tl2::MyBoxedTupleSlice& item) noexcept;
bool MyBoxedTupleSliceRead(::basictl::tl_istream & s, ::tl2::MyBoxedTupleSlice& item) noexcept; 
bool MyBoxedTupleSliceWrite(::basictl::tl_ostream & s, const ::tl2::MyBoxedTupleSlice& item) noexcept;
bool MyBoxedTupleSliceReadBoxed(::basictl::tl_istream & s, ::tl2::MyBoxedTupleSlice& item);
bool MyBoxedTupleSliceWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyBoxedTupleSlice& item);

}} // namespace tl2::details

