#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/myMaybe1.h"

namespace tl2 { namespace details { 

void MyMaybe1Reset(::tl2::MyMaybe1& item) noexcept;

bool MyMaybe1WriteJSON(std::ostream& s, const ::tl2::MyMaybe1& item) noexcept;
bool MyMaybe1Read(::basictl::tl_istream & s, ::tl2::MyMaybe1& item) noexcept; 
bool MyMaybe1Write(::basictl::tl_ostream & s, const ::tl2::MyMaybe1& item) noexcept;
bool MyMaybe1ReadBoxed(::basictl::tl_istream & s, ::tl2::MyMaybe1& item);
bool MyMaybe1WriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyMaybe1& item);

}} // namespace tl2::details

