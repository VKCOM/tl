// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service1/functions/service1.setOrIncr.h"
#include "service1/types/service1.Value.h"

namespace tl2 { namespace details { 

void Service1SetOrIncrReset(::tl2::service1::SetOrIncr& item) noexcept;

bool Service1SetOrIncrWriteJSON(std::ostream& s, const ::tl2::service1::SetOrIncr& item) noexcept;
bool Service1SetOrIncrRead(::basictl::tl_istream & s, ::tl2::service1::SetOrIncr& item) noexcept; 
bool Service1SetOrIncrWrite(::basictl::tl_ostream & s, const ::tl2::service1::SetOrIncr& item) noexcept;
bool Service1SetOrIncrReadBoxed(::basictl::tl_istream & s, ::tl2::service1::SetOrIncr& item);
bool Service1SetOrIncrWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::SetOrIncr& item);

bool Service1SetOrIncrReadResult(::basictl::tl_istream & s, ::tl2::service1::SetOrIncr& item, ::tl2::service1::Value& result);
bool Service1SetOrIncrWriteResult(::basictl::tl_ostream & s, ::tl2::service1::SetOrIncr& item, ::tl2::service1::Value& result);
		
}} // namespace tl2::details

