#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/boxedTupleSlice1.hpp"
#include "../types/int.hpp"

namespace tl2 { namespace details { 

void BoxedTupleSlice1Reset(::tl2::BoxedTupleSlice1& item);

bool BoxedTupleSlice1WriteJSON(std::ostream& s, const ::tl2::BoxedTupleSlice1& item);
bool BoxedTupleSlice1Read(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice1& item);
bool BoxedTupleSlice1Write(::basictl::tl_ostream & s, const ::tl2::BoxedTupleSlice1& item);
bool BoxedTupleSlice1ReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice1& item);
bool BoxedTupleSlice1WriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedTupleSlice1& item);

bool BoxedTupleSlice1ReadResult(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice1& item, std::vector<int32_t>& result);
bool BoxedTupleSlice1WriteResult(::basictl::tl_ostream & s, ::tl2::BoxedTupleSlice1& item, std::vector<int32_t>& result);
		
}} // namespace tl2::details

