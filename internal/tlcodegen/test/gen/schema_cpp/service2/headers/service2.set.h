#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service2/functions/service2.set.h"
#include "__common_namespace/types/true.h"

namespace tl2 { namespace details { 

void Service2SetReset(::tl2::service2::Set& item) noexcept;

bool Service2SetWriteJSON(std::ostream& s, const ::tl2::service2::Set& item) noexcept;
bool Service2SetRead(::basictl::tl_istream & s, ::tl2::service2::Set& item) noexcept; 
bool Service2SetWrite(::basictl::tl_ostream & s, const ::tl2::service2::Set& item) noexcept;
bool Service2SetReadBoxed(::basictl::tl_istream & s, ::tl2::service2::Set& item);
bool Service2SetWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service2::Set& item);

bool Service2SetReadResult(::basictl::tl_istream & s, ::tl2::service2::Set& item, ::tl2::True& result);
bool Service2SetWriteResult(::basictl::tl_ostream & s, ::tl2::service2::Set& item, ::tl2::True& result);
		
}} // namespace tl2::details

