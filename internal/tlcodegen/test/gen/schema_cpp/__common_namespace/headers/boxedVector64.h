#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/boxedVector64.h"
#include "../types/long.h"

namespace tl2 { namespace details { 

void BoxedVector64Reset(::tl2::BoxedVector64& item);

bool BoxedVector64WriteJSON(std::ostream& s, const ::tl2::BoxedVector64& item);
bool BoxedVector64Read(::basictl::tl_istream & s, ::tl2::BoxedVector64& item);
bool BoxedVector64Write(::basictl::tl_ostream & s, const ::tl2::BoxedVector64& item);
bool BoxedVector64ReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedVector64& item);
bool BoxedVector64WriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedVector64& item);

bool BoxedVector64ReadResult(::basictl::tl_istream & s, ::tl2::BoxedVector64& item, std::vector<int64_t>& result);
bool BoxedVector64WriteResult(::basictl::tl_ostream & s, ::tl2::BoxedVector64& item, std::vector<int64_t>& result);
		
}} // namespace tl2::details

