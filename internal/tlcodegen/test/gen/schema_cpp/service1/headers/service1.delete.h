#pragma once

#include "../../basictl/io_streams.h"
#include "../functions/service1.delete.h"
#include "../../__common_namespace/types/Bool.h"

namespace tl2 { namespace details { 

void Service1DeleteReset(::tl2::service1::Delete& item);

bool Service1DeleteWriteJSON(std::ostream& s, const ::tl2::service1::Delete& item);
bool Service1DeleteRead(::basictl::tl_istream & s, ::tl2::service1::Delete& item);
bool Service1DeleteWrite(::basictl::tl_ostream & s, const ::tl2::service1::Delete& item);
bool Service1DeleteReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Delete& item);
bool Service1DeleteWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Delete& item);

bool Service1DeleteReadResult(::basictl::tl_istream & s, ::tl2::service1::Delete& item, bool& result);
bool Service1DeleteWriteResult(::basictl::tl_ostream & s, ::tl2::service1::Delete& item, bool& result);
		
}} // namespace tl2::details

