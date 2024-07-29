#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../../__common/types/myDouble.hpp"

namespace tl2 { namespace details { 

void MyDoubleReset(::tl2::MyDouble& item);
bool MyDoubleRead(::basictl::tl_istream & s, ::tl2::MyDouble& item);
bool MyDoubleWrite(::basictl::tl_ostream & s, const ::tl2::MyDouble& item);
bool MyDoubleReadBoxed(::basictl::tl_istream & s, ::tl2::MyDouble& item);
bool MyDoubleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyDouble& item);

}} // namespace tl2::details

