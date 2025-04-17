#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service1/functions/service1.exists.h"
#include "__common_namespace/types/Bool.h"

namespace tl2 { namespace details { 

void Service1ExistsReset(::tl2::service1::Exists& item) noexcept;

bool Service1ExistsWriteJSON(std::ostream& s, const ::tl2::service1::Exists& item) noexcept;
bool Service1ExistsRead(::basictl::tl_istream & s, ::tl2::service1::Exists& item) noexcept; 
bool Service1ExistsWrite(::basictl::tl_ostream & s, const ::tl2::service1::Exists& item) noexcept;
bool Service1ExistsReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Exists& item);
bool Service1ExistsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Exists& item);

bool Service1ExistsReadResult(::basictl::tl_istream & s, ::tl2::service1::Exists& item, bool& result);
bool Service1ExistsWriteResult(::basictl::tl_ostream & s, ::tl2::service1::Exists& item, bool& result);
		
}} // namespace tl2::details

