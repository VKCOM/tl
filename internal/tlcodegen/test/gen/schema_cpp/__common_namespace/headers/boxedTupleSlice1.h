#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/functions/boxedTupleSlice1.h"
#include "__common_namespace/types/int.h"

namespace tl2 { namespace details { 

void BoxedTupleSlice1Reset(::tl2::BoxedTupleSlice1& item) noexcept;

bool BoxedTupleSlice1WriteJSON(std::ostream& s, const ::tl2::BoxedTupleSlice1& item) noexcept;
bool BoxedTupleSlice1Read(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice1& item) noexcept; 
bool BoxedTupleSlice1Write(::basictl::tl_ostream & s, const ::tl2::BoxedTupleSlice1& item) noexcept;
bool BoxedTupleSlice1ReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice1& item);
bool BoxedTupleSlice1WriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedTupleSlice1& item);

bool BoxedTupleSlice1ReadResult(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice1& item, std::vector<int32_t>& result);
bool BoxedTupleSlice1WriteResult(::basictl::tl_ostream & s, ::tl2::BoxedTupleSlice1& item, std::vector<int32_t>& result);
		
}} // namespace tl2::details

