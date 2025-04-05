#pragma once

#include "../../basictl/io_streams.h"
#include "../types/myMcValueVector.h"

namespace tl2 { namespace details { 

void MyMcValueVectorReset(::tl2::MyMcValueVector& item);

bool MyMcValueVectorWriteJSON(std::ostream& s, const ::tl2::MyMcValueVector& item);
bool MyMcValueVectorRead(::basictl::tl_istream & s, ::tl2::MyMcValueVector& item);
bool MyMcValueVectorWrite(::basictl::tl_ostream & s, const ::tl2::MyMcValueVector& item);
bool MyMcValueVectorReadBoxed(::basictl::tl_istream & s, ::tl2::MyMcValueVector& item);
bool MyMcValueVectorWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyMcValueVector& item);

}} // namespace tl2::details

