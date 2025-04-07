#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/myMcValueTuple.h"

namespace tl2 { namespace details { 

void MyMcValueTupleReset(::tl2::MyMcValueTuple& item);

bool MyMcValueTupleWriteJSON(std::ostream& s, const ::tl2::MyMcValueTuple& item);
bool MyMcValueTupleRead(::basictl::tl_istream & s, ::tl2::MyMcValueTuple& item);
bool MyMcValueTupleWrite(::basictl::tl_ostream & s, const ::tl2::MyMcValueTuple& item);
bool MyMcValueTupleReadBoxed(::basictl::tl_istream & s, ::tl2::MyMcValueTuple& item);
bool MyMcValueTupleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyMcValueTuple& item);

}} // namespace tl2::details

