#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../../__common/types/myMcValueTuple.hpp"

namespace tl2 { namespace details { 

void MyMcValueTupleReset(::tl2::MyMcValueTuple& item);
bool MyMcValueTupleRead(::basictl::tl_istream & s, ::tl2::MyMcValueTuple& item);
bool MyMcValueTupleWrite(::basictl::tl_ostream & s, const ::tl2::MyMcValueTuple& item);
bool MyMcValueTupleReadBoxed(::basictl::tl_istream & s, ::tl2::MyMcValueTuple& item);
bool MyMcValueTupleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyMcValueTuple& item);

}} // namespace tl2::details

