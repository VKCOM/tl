// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service1/functions/service1.replaceOrIncr.h"
#include "service1/types/service1.Value.h"

namespace tl2 { namespace details { 

void Service1ReplaceOrIncrReset(::tl2::service1::ReplaceOrIncr& item) noexcept;

bool Service1ReplaceOrIncrWriteJSON(std::ostream& s, const ::tl2::service1::ReplaceOrIncr& item) noexcept;
bool Service1ReplaceOrIncrRead(::basictl::tl_istream & s, ::tl2::service1::ReplaceOrIncr& item) noexcept; 
bool Service1ReplaceOrIncrWrite(::basictl::tl_ostream & s, const ::tl2::service1::ReplaceOrIncr& item) noexcept;
bool Service1ReplaceOrIncrReadBoxed(::basictl::tl_istream & s, ::tl2::service1::ReplaceOrIncr& item);
bool Service1ReplaceOrIncrWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::ReplaceOrIncr& item);

bool Service1ReplaceOrIncrReadResult(::basictl::tl_istream & s, ::tl2::service1::ReplaceOrIncr& item, ::tl2::service1::Value& result);
bool Service1ReplaceOrIncrWriteResult(::basictl::tl_ostream & s, ::tl2::service1::ReplaceOrIncr& item, ::tl2::service1::Value& result);
		
}} // namespace tl2::details

