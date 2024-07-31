#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/boxedInt.hpp"

namespace tl2 { namespace details { 

void BoxedIntReset(::tl2::BoxedInt& item);

bool BoxedIntWriteJSON(std::ostream& s, const ::tl2::BoxedInt& item);
bool BoxedIntRead(::basictl::tl_istream & s, ::tl2::BoxedInt& item);
bool BoxedIntWrite(::basictl::tl_ostream & s, const ::tl2::BoxedInt& item);
bool BoxedIntReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedInt& item);
bool BoxedIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedInt& item);

bool BoxedIntReadResult(::basictl::tl_istream & s, ::tl2::BoxedInt& item, int32_t& result);
bool BoxedIntWriteResult(::basictl::tl_ostream & s, ::tl2::BoxedInt& item, int32_t& result);
		
}} // namespace tl2::details

