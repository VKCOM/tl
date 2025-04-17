#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/myMaybe2.h"

namespace tl2 { namespace details { 

void MyMaybe2Reset(::tl2::MyMaybe2& item) noexcept;

bool MyMaybe2WriteJSON(std::ostream& s, const ::tl2::MyMaybe2& item) noexcept;
bool MyMaybe2Read(::basictl::tl_istream & s, ::tl2::MyMaybe2& item) noexcept; 
bool MyMaybe2Write(::basictl::tl_ostream & s, const ::tl2::MyMaybe2& item) noexcept;
bool MyMaybe2ReadBoxed(::basictl::tl_istream & s, ::tl2::MyMaybe2& item);
bool MyMaybe2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyMaybe2& item);

}} // namespace tl2::details

