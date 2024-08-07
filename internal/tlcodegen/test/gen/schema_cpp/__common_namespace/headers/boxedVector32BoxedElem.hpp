#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/boxedVector32BoxedElem.hpp"
#include "../types/int.hpp"

namespace tl2 { namespace details { 

void BoxedVector32BoxedElemReset(::tl2::BoxedVector32BoxedElem& item);

bool BoxedVector32BoxedElemWriteJSON(std::ostream& s, const ::tl2::BoxedVector32BoxedElem& item);
bool BoxedVector32BoxedElemRead(::basictl::tl_istream & s, ::tl2::BoxedVector32BoxedElem& item);
bool BoxedVector32BoxedElemWrite(::basictl::tl_ostream & s, const ::tl2::BoxedVector32BoxedElem& item);
bool BoxedVector32BoxedElemReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedVector32BoxedElem& item);
bool BoxedVector32BoxedElemWriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedVector32BoxedElem& item);

bool BoxedVector32BoxedElemReadResult(::basictl::tl_istream & s, ::tl2::BoxedVector32BoxedElem& item, std::vector<int32_t>& result);
bool BoxedVector32BoxedElemWriteResult(::basictl::tl_ostream & s, ::tl2::BoxedVector32BoxedElem& item, std::vector<int32_t>& result);
		
}} // namespace tl2::details

