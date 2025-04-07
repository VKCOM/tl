#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/getMyDictOfInt.h"
#include "../types/myDictOfInt.h"

namespace tl2 { namespace details { 

void GetMyDictOfIntReset(::tl2::GetMyDictOfInt& item);

bool GetMyDictOfIntWriteJSON(std::ostream& s, const ::tl2::GetMyDictOfInt& item);
bool GetMyDictOfIntRead(::basictl::tl_istream & s, ::tl2::GetMyDictOfInt& item);
bool GetMyDictOfIntWrite(::basictl::tl_ostream & s, const ::tl2::GetMyDictOfInt& item);
bool GetMyDictOfIntReadBoxed(::basictl::tl_istream & s, ::tl2::GetMyDictOfInt& item);
bool GetMyDictOfIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetMyDictOfInt& item);

bool GetMyDictOfIntReadResult(::basictl::tl_istream & s, ::tl2::GetMyDictOfInt& item, ::tl2::MyDictOfInt& result);
bool GetMyDictOfIntWriteResult(::basictl::tl_ostream & s, ::tl2::GetMyDictOfInt& item, ::tl2::MyDictOfInt& result);
		
}} // namespace tl2::details

