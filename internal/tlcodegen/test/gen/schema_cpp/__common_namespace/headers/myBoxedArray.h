#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/myBoxedArray.h"

namespace tl2 { namespace details { 

void MyBoxedArrayReset(::tl2::MyBoxedArray& item) noexcept;

bool MyBoxedArrayWriteJSON(std::ostream& s, const ::tl2::MyBoxedArray& item) noexcept;
bool MyBoxedArrayRead(::basictl::tl_istream & s, ::tl2::MyBoxedArray& item) noexcept; 
bool MyBoxedArrayWrite(::basictl::tl_ostream & s, const ::tl2::MyBoxedArray& item) noexcept;
bool MyBoxedArrayReadBoxed(::basictl::tl_istream & s, ::tl2::MyBoxedArray& item);
bool MyBoxedArrayWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyBoxedArray& item);

}} // namespace tl2::details

