#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/myBoxedVectorSlice.h"

namespace tl2 { namespace details { 

void MyBoxedVectorSliceReset(::tl2::MyBoxedVectorSlice& item) noexcept;

bool MyBoxedVectorSliceWriteJSON(std::ostream& s, const ::tl2::MyBoxedVectorSlice& item) noexcept;
bool MyBoxedVectorSliceRead(::basictl::tl_istream & s, ::tl2::MyBoxedVectorSlice& item) noexcept; 
bool MyBoxedVectorSliceWrite(::basictl::tl_ostream & s, const ::tl2::MyBoxedVectorSlice& item) noexcept;
bool MyBoxedVectorSliceReadBoxed(::basictl::tl_istream & s, ::tl2::MyBoxedVectorSlice& item);
bool MyBoxedVectorSliceWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyBoxedVectorSlice& item);

}} // namespace tl2::details

