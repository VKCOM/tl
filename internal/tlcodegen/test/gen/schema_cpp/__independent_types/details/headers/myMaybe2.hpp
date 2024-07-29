#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../../__common/types/myMaybe2.hpp"

namespace tl2 { namespace details { 

void MyMaybe2Reset(::tl2::MyMaybe2& item);
bool MyMaybe2Read(::basictl::tl_istream & s, ::tl2::MyMaybe2& item);
bool MyMaybe2Write(::basictl::tl_ostream & s, const ::tl2::MyMaybe2& item);
bool MyMaybe2ReadBoxed(::basictl::tl_istream & s, ::tl2::MyMaybe2& item);
bool MyMaybe2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyMaybe2& item);

}} // namespace tl2::details

