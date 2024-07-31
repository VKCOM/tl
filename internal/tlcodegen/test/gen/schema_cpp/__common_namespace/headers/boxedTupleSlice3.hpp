#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/boxedTupleSlice3.hpp"
#include "../types/int.hpp"

namespace tl2 { namespace details { 

void BoxedTupleSlice3Reset(::tl2::BoxedTupleSlice3& item);

bool BoxedTupleSlice3WriteJSON(std::ostream& s, const ::tl2::BoxedTupleSlice3& item);
bool BoxedTupleSlice3Read(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice3& item);
bool BoxedTupleSlice3Write(::basictl::tl_ostream & s, const ::tl2::BoxedTupleSlice3& item);
bool BoxedTupleSlice3ReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice3& item);
bool BoxedTupleSlice3WriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedTupleSlice3& item);

bool BoxedTupleSlice3ReadResult(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice3& item, std::vector<int32_t>& result);
bool BoxedTupleSlice3WriteResult(::basictl::tl_ostream & s, ::tl2::BoxedTupleSlice3& item, std::vector<int32_t>& result);
		
}} // namespace tl2::details

