// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service3/functions/service3.getLimits.h"
#include "service3/types/service3.limits.h"

namespace tlgen { namespace details { 

void Service3GetLimitsReset(::tlgen::service3::GetLimits& item) noexcept;

bool Service3GetLimitsWriteJSON(std::ostream& s, const ::tlgen::service3::GetLimits& item) noexcept;
bool Service3GetLimitsRead(::tlgen::basictl::tl_istream & s, ::tlgen::service3::GetLimits& item) noexcept; 
bool Service3GetLimitsWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::service3::GetLimits& item) noexcept;
bool Service3GetLimitsReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::service3::GetLimits& item);
bool Service3GetLimitsWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::service3::GetLimits& item);

bool Service3GetLimitsReadResult(::tlgen::basictl::tl_istream & s, const ::tlgen::service3::GetLimits& item, ::tlgen::service3::Limits& result);
bool Service3GetLimitsWriteResult(::tlgen::basictl::tl_ostream & s, const ::tlgen::service3::GetLimits& item, const ::tlgen::service3::Limits& result);
    
}} // namespace tlgen::details

