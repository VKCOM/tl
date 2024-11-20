#pragma once

#include "../../../basics/basictl.h"
#include "../types/service1.not_found.h"

namespace tl2 { namespace details { 

void Service1NotFoundReset(::tl2::service1::Not_found& item);

bool Service1NotFoundWriteJSON(std::ostream& s, const ::tl2::service1::Not_found& item);
bool Service1NotFoundRead(::basictl::tl_istream & s, ::tl2::service1::Not_found& item);
bool Service1NotFoundWrite(::basictl::tl_ostream & s, const ::tl2::service1::Not_found& item);
bool Service1NotFoundReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Not_found& item);
bool Service1NotFoundWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Not_found& item);

}} // namespace tl2::details

