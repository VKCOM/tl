#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/service2.objectId.h"

namespace tl2 { namespace details { 

void Service2ObjectIdReset(::tl2::service2::ObjectId& item) noexcept;

bool Service2ObjectIdWriteJSON(std::ostream& s, const ::tl2::service2::ObjectId& item, uint32_t nat_objectIdLength) noexcept;
bool Service2ObjectIdRead(::basictl::tl_istream & s, ::tl2::service2::ObjectId& item, uint32_t nat_objectIdLength) noexcept; 
bool Service2ObjectIdWrite(::basictl::tl_ostream & s, const ::tl2::service2::ObjectId& item, uint32_t nat_objectIdLength) noexcept;
bool Service2ObjectIdReadBoxed(::basictl::tl_istream & s, ::tl2::service2::ObjectId& item, uint32_t nat_objectIdLength);
bool Service2ObjectIdWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service2::ObjectId& item, uint32_t nat_objectIdLength);

}} // namespace tl2::details

