// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service1/functions/service1.disableKeysStat.h"
#include "__common_namespace/types/Bool.h"

namespace tlgen { namespace details { 

void Service1DisableKeysStatReset(::tlgen::service1::DisableKeysStat& item) noexcept;

bool Service1DisableKeysStatWriteJSON(std::ostream& s, const ::tlgen::service1::DisableKeysStat& item) noexcept;
bool Service1DisableKeysStatRead(::tlgen::basictl::tl_istream & s, ::tlgen::service1::DisableKeysStat& item) noexcept; 
bool Service1DisableKeysStatWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::service1::DisableKeysStat& item) noexcept;
bool Service1DisableKeysStatReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::service1::DisableKeysStat& item);
bool Service1DisableKeysStatWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::service1::DisableKeysStat& item);

bool Service1DisableKeysStatReadResult(::tlgen::basictl::tl_istream & s, const ::tlgen::service1::DisableKeysStat& item, bool& result);
bool Service1DisableKeysStatWriteResult(::tlgen::basictl::tl_ostream & s, const ::tlgen::service1::DisableKeysStat& item, const bool& result);
    
}} // namespace tlgen::details

