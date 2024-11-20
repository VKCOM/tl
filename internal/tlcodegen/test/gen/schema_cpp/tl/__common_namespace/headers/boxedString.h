#pragma once

#include "../../../basics/basictl.h"
#include "../functions/boxedString.h"

namespace tl2 { namespace details { 

void BoxedStringReset(::tl2::BoxedString& item);

bool BoxedStringWriteJSON(std::ostream& s, const ::tl2::BoxedString& item);
bool BoxedStringRead(::basictl::tl_istream & s, ::tl2::BoxedString& item);
bool BoxedStringWrite(::basictl::tl_ostream & s, const ::tl2::BoxedString& item);
bool BoxedStringReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedString& item);
bool BoxedStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedString& item);

bool BoxedStringReadResult(::basictl::tl_istream & s, ::tl2::BoxedString& item, std::string& result);
bool BoxedStringWriteResult(::basictl::tl_ostream & s, ::tl2::BoxedString& item, std::string& result);
		
}} // namespace tl2::details

