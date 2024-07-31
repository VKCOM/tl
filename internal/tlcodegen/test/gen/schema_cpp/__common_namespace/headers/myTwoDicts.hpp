#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/myTwoDicts.hpp"

namespace tl2 { namespace details { 

void MyTwoDictsReset(::tl2::MyTwoDicts& item);

bool MyTwoDictsWriteJSON(std::ostream& s, const ::tl2::MyTwoDicts& item);
bool MyTwoDictsRead(::basictl::tl_istream & s, ::tl2::MyTwoDicts& item);
bool MyTwoDictsWrite(::basictl::tl_ostream & s, const ::tl2::MyTwoDicts& item);
bool MyTwoDictsReadBoxed(::basictl::tl_istream & s, ::tl2::MyTwoDicts& item);
bool MyTwoDictsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyTwoDicts& item);

}} // namespace tl2::details

