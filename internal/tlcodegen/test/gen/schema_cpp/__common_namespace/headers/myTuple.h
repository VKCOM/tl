#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/myTuple.h"

namespace tl2 { namespace details { 

void MyTuple10Reset(::tl2::MyTuplen<10>& item);

bool MyTuple10WriteJSON(std::ostream& s, const ::tl2::MyTuplen<10>& item);
bool MyTuple10Read(::basictl::tl_istream & s, ::tl2::MyTuplen<10>& item);
bool MyTuple10Write(::basictl::tl_ostream & s, const ::tl2::MyTuplen<10>& item);
bool MyTuple10ReadBoxed(::basictl::tl_istream & s, ::tl2::MyTuplen<10>& item);
bool MyTuple10WriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyTuplen<10>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

bool MyTuple10MaybeWriteJSON(std::ostream & s, const std::optional<::tl2::MyTuplen<10>>& item);

bool MyTuple10MaybeReadBoxed(::basictl::tl_istream & s, std::optional<::tl2::MyTuplen<10>>& item);
bool MyTuple10MaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<::tl2::MyTuplen<10>>& item);


}} // namespace tl2::details

