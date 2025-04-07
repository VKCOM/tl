#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/getMyValue.h"
#include "../types/MyValue.h"

namespace tl2 { namespace details { 

void GetMyValueReset(::tl2::GetMyValue& item);

bool GetMyValueWriteJSON(std::ostream& s, const ::tl2::GetMyValue& item);
bool GetMyValueRead(::basictl::tl_istream & s, ::tl2::GetMyValue& item);
bool GetMyValueWrite(::basictl::tl_ostream & s, const ::tl2::GetMyValue& item);
bool GetMyValueReadBoxed(::basictl::tl_istream & s, ::tl2::GetMyValue& item);
bool GetMyValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetMyValue& item);

bool GetMyValueReadResult(::basictl::tl_istream & s, ::tl2::GetMyValue& item, ::tl2::MyValue& result);
bool GetMyValueWriteResult(::basictl::tl_ostream & s, ::tl2::GetMyValue& item, ::tl2::MyValue& result);
		
}} // namespace tl2::details

