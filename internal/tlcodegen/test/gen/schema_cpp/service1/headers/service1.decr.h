#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service1/functions/service1.decr.h"
#include "service1/types/service1.Value.h"

namespace tl2 { namespace details { 

void Service1DecrReset(::tl2::service1::Decr& item) noexcept;

bool Service1DecrWriteJSON(std::ostream& s, const ::tl2::service1::Decr& item) noexcept;
bool Service1DecrRead(::basictl::tl_istream & s, ::tl2::service1::Decr& item) noexcept; 
bool Service1DecrWrite(::basictl::tl_ostream & s, const ::tl2::service1::Decr& item) noexcept;
bool Service1DecrReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Decr& item);
bool Service1DecrWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Decr& item);

bool Service1DecrReadResult(::basictl::tl_istream & s, ::tl2::service1::Decr& item, ::tl2::service1::Value& result);
bool Service1DecrWriteResult(::basictl::tl_ostream & s, ::tl2::service1::Decr& item, ::tl2::service1::Value& result);
		
}} // namespace tl2::details

