#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/functions/boxedInt.h"

namespace tl2 { namespace details { 

void BoxedIntReset(::tl2::BoxedInt& item) noexcept;

bool BoxedIntWriteJSON(std::ostream& s, const ::tl2::BoxedInt& item) noexcept;
bool BoxedIntRead(::basictl::tl_istream & s, ::tl2::BoxedInt& item) noexcept; 
bool BoxedIntWrite(::basictl::tl_ostream & s, const ::tl2::BoxedInt& item) noexcept;
bool BoxedIntReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedInt& item);
bool BoxedIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedInt& item);

bool BoxedIntReadResult(::basictl::tl_istream & s, ::tl2::BoxedInt& item, int32_t& result);
bool BoxedIntWriteResult(::basictl::tl_ostream & s, ::tl2::BoxedInt& item, int32_t& result);
		
}} // namespace tl2::details

