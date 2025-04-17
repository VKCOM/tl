// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service1/functions/service1.get.h"
#include "service1/types/service1.Value.h"

namespace tl2 { namespace details { 

void Service1GetReset(::tl2::service1::Get& item) noexcept;

bool Service1GetWriteJSON(std::ostream& s, const ::tl2::service1::Get& item) noexcept;
bool Service1GetRead(::basictl::tl_istream & s, ::tl2::service1::Get& item) noexcept; 
bool Service1GetWrite(::basictl::tl_ostream & s, const ::tl2::service1::Get& item) noexcept;
bool Service1GetReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Get& item);
bool Service1GetWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Get& item);

bool Service1GetReadResult(::basictl::tl_istream & s, ::tl2::service1::Get& item, ::tl2::service1::Value& result);
bool Service1GetWriteResult(::basictl::tl_ostream & s, ::tl2::service1::Get& item, ::tl2::service1::Value& result);
		
}} // namespace tl2::details

