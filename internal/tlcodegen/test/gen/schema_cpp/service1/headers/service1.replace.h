#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/service1.replace.h"
#include "../../__common_namespace/types/Bool.h"

namespace tl2 { namespace details { 

void Service1ReplaceReset(::tl2::service1::Replace& item) noexcept;

bool Service1ReplaceWriteJSON(std::ostream& s, const ::tl2::service1::Replace& item) noexcept;
bool Service1ReplaceRead(::basictl::tl_istream & s, ::tl2::service1::Replace& item) noexcept; 
bool Service1ReplaceWrite(::basictl::tl_ostream & s, const ::tl2::service1::Replace& item) noexcept;
bool Service1ReplaceReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Replace& item);
bool Service1ReplaceWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Replace& item);

bool Service1ReplaceReadResult(::basictl::tl_istream & s, ::tl2::service1::Replace& item, bool& result);
bool Service1ReplaceWriteResult(::basictl::tl_ostream & s, ::tl2::service1::Replace& item, bool& result);
		
}} // namespace tl2::details

