#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/functions/boxedArray.h"
#include "__common_namespace/types/myBoxedArray.h"

namespace tl2 { namespace details { 

void BoxedArrayReset(::tl2::BoxedArray& item) noexcept;

bool BoxedArrayWriteJSON(std::ostream& s, const ::tl2::BoxedArray& item) noexcept;
bool BoxedArrayRead(::basictl::tl_istream & s, ::tl2::BoxedArray& item) noexcept; 
bool BoxedArrayWrite(::basictl::tl_ostream & s, const ::tl2::BoxedArray& item) noexcept;
bool BoxedArrayReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedArray& item);
bool BoxedArrayWriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedArray& item);

bool BoxedArrayReadResult(::basictl::tl_istream & s, ::tl2::BoxedArray& item, ::tl2::MyBoxedArray& result);
bool BoxedArrayWriteResult(::basictl::tl_ostream & s, ::tl2::BoxedArray& item, ::tl2::MyBoxedArray& result);
		
}} // namespace tl2::details

