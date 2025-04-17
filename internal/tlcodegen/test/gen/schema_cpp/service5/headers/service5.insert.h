#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service5/functions/service5.insert.h"
#include "service5/types/service5.Output.h"

namespace tl2 { namespace details { 

void Service5InsertReset(::tl2::service5::Insert& item) noexcept;

bool Service5InsertWriteJSON(std::ostream& s, const ::tl2::service5::Insert& item) noexcept;
bool Service5InsertRead(::basictl::tl_istream & s, ::tl2::service5::Insert& item) noexcept; 
bool Service5InsertWrite(::basictl::tl_ostream & s, const ::tl2::service5::Insert& item) noexcept;
bool Service5InsertReadBoxed(::basictl::tl_istream & s, ::tl2::service5::Insert& item);
bool Service5InsertWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service5::Insert& item);

bool Service5InsertReadResult(::basictl::tl_istream & s, ::tl2::service5::Insert& item, ::tl2::service5::Output& result);
bool Service5InsertWriteResult(::basictl::tl_ostream & s, ::tl2::service5::Insert& item, ::tl2::service5::Output& result);
		
}} // namespace tl2::details

