#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/boxedVector32.hpp"
#include "../types/int.hpp"

namespace tl2 { namespace details { 

void BoxedVector32Reset(::tl2::BoxedVector32& item);

bool BoxedVector32WriteJSON(std::ostream& s, const ::tl2::BoxedVector32& item);
bool BoxedVector32Read(::basictl::tl_istream & s, ::tl2::BoxedVector32& item);
bool BoxedVector32Write(::basictl::tl_ostream & s, const ::tl2::BoxedVector32& item);
bool BoxedVector32ReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedVector32& item);
bool BoxedVector32WriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedVector32& item);

bool BoxedVector32ReadResult(::basictl::tl_istream & s, ::tl2::BoxedVector32& item, std::vector<int32_t>& result);
bool BoxedVector32WriteResult(::basictl::tl_ostream & s, ::tl2::BoxedVector32& item, std::vector<int32_t>& result);
		
}} // namespace tl2::details

