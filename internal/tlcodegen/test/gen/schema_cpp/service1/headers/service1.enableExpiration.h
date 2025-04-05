#pragma once

#include "../../basictl/io_streams.h"
#include "../functions/service1.enableExpiration.h"
#include "../../__common_namespace/types/Bool.h"

namespace tl2 { namespace details { 

void Service1EnableExpirationReset(::tl2::service1::EnableExpiration& item);

bool Service1EnableExpirationWriteJSON(std::ostream& s, const ::tl2::service1::EnableExpiration& item);
bool Service1EnableExpirationRead(::basictl::tl_istream & s, ::tl2::service1::EnableExpiration& item);
bool Service1EnableExpirationWrite(::basictl::tl_ostream & s, const ::tl2::service1::EnableExpiration& item);
bool Service1EnableExpirationReadBoxed(::basictl::tl_istream & s, ::tl2::service1::EnableExpiration& item);
bool Service1EnableExpirationWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::EnableExpiration& item);

bool Service1EnableExpirationReadResult(::basictl::tl_istream & s, ::tl2::service1::EnableExpiration& item, bool& result);
bool Service1EnableExpirationWriteResult(::basictl::tl_ostream & s, ::tl2::service1::EnableExpiration& item, bool& result);
		
}} // namespace tl2::details

