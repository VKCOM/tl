#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../../__common/types/myBoxedVectorSlice.hpp"

namespace tl2 { namespace details { 

void MyBoxedVectorSliceReset(::tl2::MyBoxedVectorSlice& item);
bool MyBoxedVectorSliceRead(::basictl::tl_istream & s, ::tl2::MyBoxedVectorSlice& item);
bool MyBoxedVectorSliceWrite(::basictl::tl_ostream & s, const ::tl2::MyBoxedVectorSlice& item);
bool MyBoxedVectorSliceReadBoxed(::basictl::tl_istream & s, ::tl2::MyBoxedVectorSlice& item);
bool MyBoxedVectorSliceWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyBoxedVectorSlice& item);

}} // namespace tl2::details

