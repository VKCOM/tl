#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/service1.append.hpp"
#include "../../__common_namespace/types/Bool.hpp"

namespace tl2 { namespace details { 

void Service1AppendReset(::tl2::service1::Append& item);
bool Service1AppendRead(::basictl::tl_istream & s, ::tl2::service1::Append& item);
bool Service1AppendWrite(::basictl::tl_ostream & s, const ::tl2::service1::Append& item);
bool Service1AppendReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Append& item);
bool Service1AppendWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Append& item);

bool Service1AppendReadResult(::basictl::tl_istream & s, ::tl2::service1::Append& item, bool& result);
bool Service1AppendWriteResult(::basictl::tl_ostream & s, ::tl2::service1::Append& item, bool& result);
		
}} // namespace tl2::details

