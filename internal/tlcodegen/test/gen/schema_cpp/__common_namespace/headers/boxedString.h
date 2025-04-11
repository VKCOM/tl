#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/boxedString.h"

namespace tl2 { namespace details { 

void BoxedStringReset(::tl2::BoxedString& item) noexcept;

bool BoxedStringWriteJSON(std::ostream& s, const ::tl2::BoxedString& item) noexcept;
bool BoxedStringRead(::basictl::tl_istream & s, ::tl2::BoxedString& item) noexcept; 
bool BoxedStringWrite(::basictl::tl_ostream & s, const ::tl2::BoxedString& item) noexcept;
bool BoxedStringReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedString& item);
bool BoxedStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedString& item);

bool BoxedStringReadResult(::basictl::tl_istream & s, ::tl2::BoxedString& item, std::string& result);
bool BoxedStringWriteResult(::basictl::tl_ostream & s, ::tl2::BoxedString& item, std::string& result);
		
}} // namespace tl2::details

