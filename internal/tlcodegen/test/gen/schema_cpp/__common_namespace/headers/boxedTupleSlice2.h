// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/functions/boxedTupleSlice2.h"
#include "__common_namespace/types/myBoxedTupleSlice.h"

namespace tl2 { namespace details { 

void BoxedTupleSlice2Reset(::tl2::BoxedTupleSlice2& item) noexcept;

bool BoxedTupleSlice2WriteJSON(std::ostream& s, const ::tl2::BoxedTupleSlice2& item) noexcept;
bool BoxedTupleSlice2Read(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice2& item) noexcept; 
bool BoxedTupleSlice2Write(::basictl::tl_ostream & s, const ::tl2::BoxedTupleSlice2& item) noexcept;
bool BoxedTupleSlice2ReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice2& item);
bool BoxedTupleSlice2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedTupleSlice2& item);

bool BoxedTupleSlice2ReadResult(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice2& item, ::tl2::MyBoxedTupleSlice& result);
bool BoxedTupleSlice2WriteResult(::basictl::tl_ostream & s, ::tl2::BoxedTupleSlice2& item, ::tl2::MyBoxedTupleSlice& result);
		
}} // namespace tl2::details

