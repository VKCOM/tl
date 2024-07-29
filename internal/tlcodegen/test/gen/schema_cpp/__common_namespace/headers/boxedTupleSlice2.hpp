#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/boxedTupleSlice2.hpp"
#include "../types/myBoxedTupleSlice.hpp"

namespace tl2 { namespace details { 

void BoxedTupleSlice2Reset(::tl2::BoxedTupleSlice2& item);
bool BoxedTupleSlice2Read(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice2& item);
bool BoxedTupleSlice2Write(::basictl::tl_ostream & s, const ::tl2::BoxedTupleSlice2& item);
bool BoxedTupleSlice2ReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice2& item);
bool BoxedTupleSlice2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedTupleSlice2& item);

bool BoxedTupleSlice2ReadResult(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice2& item, ::tl2::MyBoxedTupleSlice& result);
bool BoxedTupleSlice2WriteResult(::basictl::tl_ostream & s, ::tl2::BoxedTupleSlice2& item, ::tl2::MyBoxedTupleSlice& result);
		
}} // namespace tl2::details

