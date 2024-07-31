#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/service1.cas.hpp"
#include "../../__common_namespace/types/Bool.hpp"

namespace tl2 { namespace details { 

void Service1CasReset(::tl2::service1::Cas& item);

bool Service1CasWriteJSON(std::ostream& s, const ::tl2::service1::Cas& item);
bool Service1CasRead(::basictl::tl_istream & s, ::tl2::service1::Cas& item);
bool Service1CasWrite(::basictl::tl_ostream & s, const ::tl2::service1::Cas& item);
bool Service1CasReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Cas& item);
bool Service1CasWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Cas& item);

bool Service1CasReadResult(::basictl::tl_istream & s, ::tl2::service1::Cas& item, bool& result);
bool Service1CasWriteResult(::basictl::tl_ostream & s, ::tl2::service1::Cas& item, bool& result);
		
}} // namespace tl2::details

