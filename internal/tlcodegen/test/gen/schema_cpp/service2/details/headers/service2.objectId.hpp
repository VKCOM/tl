#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../types/service2.objectId.hpp"

namespace tl2 { namespace details { 

void Service2ObjectIdReset(::tl2::service2::ObjectId& item);
bool Service2ObjectIdRead(::basictl::tl_istream & s, ::tl2::service2::ObjectId& item, uint32_t nat_objectIdLength);
bool Service2ObjectIdWrite(::basictl::tl_ostream & s, const ::tl2::service2::ObjectId& item, uint32_t nat_objectIdLength);
bool Service2ObjectIdReadBoxed(::basictl::tl_istream & s, ::tl2::service2::ObjectId& item, uint32_t nat_objectIdLength);
bool Service2ObjectIdWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service2::ObjectId& item, uint32_t nat_objectIdLength);

}} // namespace tl2::details

