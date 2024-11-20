#pragma once

#include "../../../basics/basictl.h"
#include "../functions/boxedArray.h"
#include "../types/myBoxedArray.h"

namespace tl2 { namespace details { 

void BoxedArrayReset(::tl2::BoxedArray& item);

bool BoxedArrayWriteJSON(std::ostream& s, const ::tl2::BoxedArray& item);
bool BoxedArrayRead(::basictl::tl_istream & s, ::tl2::BoxedArray& item);
bool BoxedArrayWrite(::basictl::tl_ostream & s, const ::tl2::BoxedArray& item);
bool BoxedArrayReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedArray& item);
bool BoxedArrayWriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedArray& item);

bool BoxedArrayReadResult(::basictl::tl_istream & s, ::tl2::BoxedArray& item, ::tl2::MyBoxedArray& result);
bool BoxedArrayWriteResult(::basictl::tl_ostream & s, ::tl2::BoxedArray& item, ::tl2::MyBoxedArray& result);
		
}} // namespace tl2::details

