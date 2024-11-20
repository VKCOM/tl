#pragma once

#include "../../../basics/basictl.h"
#include "../types/service4.object.h"

namespace tl2 { namespace details { 

void Service4ObjectReset(::tl2::service4::Object& item);

bool Service4ObjectWriteJSON(std::ostream& s, const ::tl2::service4::Object& item);
bool Service4ObjectRead(::basictl::tl_istream & s, ::tl2::service4::Object& item);
bool Service4ObjectWrite(::basictl::tl_ostream & s, const ::tl2::service4::Object& item);
bool Service4ObjectReadBoxed(::basictl::tl_istream & s, ::tl2::service4::Object& item);
bool Service4ObjectWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service4::Object& item);

}} // namespace tl2::details

