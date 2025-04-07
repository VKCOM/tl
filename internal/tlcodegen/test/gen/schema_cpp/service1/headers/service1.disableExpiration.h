#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/service1.disableExpiration.h"
#include "../../__common_namespace/types/Bool.h"

namespace tl2 { namespace details { 

void Service1DisableExpirationReset(::tl2::service1::DisableExpiration& item);

bool Service1DisableExpirationWriteJSON(std::ostream& s, const ::tl2::service1::DisableExpiration& item);
bool Service1DisableExpirationRead(::basictl::tl_istream & s, ::tl2::service1::DisableExpiration& item);
bool Service1DisableExpirationWrite(::basictl::tl_ostream & s, const ::tl2::service1::DisableExpiration& item);
bool Service1DisableExpirationReadBoxed(::basictl::tl_istream & s, ::tl2::service1::DisableExpiration& item);
bool Service1DisableExpirationWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::DisableExpiration& item);

bool Service1DisableExpirationReadResult(::basictl::tl_istream & s, ::tl2::service1::DisableExpiration& item, bool& result);
bool Service1DisableExpirationWriteResult(::basictl::tl_ostream & s, ::tl2::service1::DisableExpiration& item, bool& result);
		
}} // namespace tl2::details

