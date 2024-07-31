#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/myBoxedArray.hpp"

namespace tl2 { namespace details { 

void MyBoxedArrayReset(::tl2::MyBoxedArray& item);

bool MyBoxedArrayWriteJSON(std::ostream& s, const ::tl2::MyBoxedArray& item);
bool MyBoxedArrayRead(::basictl::tl_istream & s, ::tl2::MyBoxedArray& item);
bool MyBoxedArrayWrite(::basictl::tl_ostream & s, const ::tl2::MyBoxedArray& item);
bool MyBoxedArrayReadBoxed(::basictl::tl_istream & s, ::tl2::MyBoxedArray& item);
bool MyBoxedArrayWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyBoxedArray& item);

}} // namespace tl2::details

