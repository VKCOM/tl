#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/service2.setObjectTtl.hpp"
#include "../../__common_namespace/types/true.hpp"

namespace tl2 { namespace details { 

void Service2SetObjectTtlReset(::tl2::service2::SetObjectTtl& item);

bool Service2SetObjectTtlWriteJSON(std::ostream& s, const ::tl2::service2::SetObjectTtl& item);
bool Service2SetObjectTtlRead(::basictl::tl_istream & s, ::tl2::service2::SetObjectTtl& item);
bool Service2SetObjectTtlWrite(::basictl::tl_ostream & s, const ::tl2::service2::SetObjectTtl& item);
bool Service2SetObjectTtlReadBoxed(::basictl::tl_istream & s, ::tl2::service2::SetObjectTtl& item);
bool Service2SetObjectTtlWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service2::SetObjectTtl& item);

bool Service2SetObjectTtlReadResult(::basictl::tl_istream & s, ::tl2::service2::SetObjectTtl& item, ::tl2::True& result);
bool Service2SetObjectTtlWriteResult(::basictl::tl_ostream & s, ::tl2::service2::SetObjectTtl& item, ::tl2::True& result);
		
}} // namespace tl2::details

