#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/myTwoDicts.h"

namespace tl2 { namespace details { 

void MyTwoDictsReset(::tl2::MyTwoDicts& item) noexcept;

bool MyTwoDictsWriteJSON(std::ostream& s, const ::tl2::MyTwoDicts& item) noexcept;
bool MyTwoDictsRead(::basictl::tl_istream & s, ::tl2::MyTwoDicts& item) noexcept; 
bool MyTwoDictsWrite(::basictl::tl_ostream & s, const ::tl2::MyTwoDicts& item) noexcept;
bool MyTwoDictsReadBoxed(::basictl::tl_istream & s, ::tl2::MyTwoDicts& item);
bool MyTwoDictsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyTwoDicts& item);

}} // namespace tl2::details

