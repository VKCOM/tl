#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/myDouble.h"

namespace tl2 { namespace details { 

void MyDoubleReset(::tl2::MyDouble& item);

bool MyDoubleWriteJSON(std::ostream& s, const ::tl2::MyDouble& item);
bool MyDoubleRead(::basictl::tl_istream & s, ::tl2::MyDouble& item);
bool MyDoubleWrite(::basictl::tl_ostream & s, const ::tl2::MyDouble& item);
bool MyDoubleReadBoxed(::basictl::tl_istream & s, ::tl2::MyDouble& item);
bool MyDoubleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyDouble& item);

}} // namespace tl2::details

