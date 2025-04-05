#pragma once

#include "../../basictl/io_streams.h"
#include "../functions/service1.cas.h"
#include "../../__common_namespace/types/Bool.h"

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

