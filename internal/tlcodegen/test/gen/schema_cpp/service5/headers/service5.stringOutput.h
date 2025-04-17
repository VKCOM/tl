#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service5/types/service5.stringOutput.h"

namespace tl2 { namespace details { 

void Service5StringOutputReset(::tl2::service5::StringOutput& item) noexcept;

bool Service5StringOutputWriteJSON(std::ostream& s, const ::tl2::service5::StringOutput& item) noexcept;
bool Service5StringOutputRead(::basictl::tl_istream & s, ::tl2::service5::StringOutput& item) noexcept; 
bool Service5StringOutputWrite(::basictl::tl_ostream & s, const ::tl2::service5::StringOutput& item) noexcept;
bool Service5StringOutputReadBoxed(::basictl::tl_istream & s, ::tl2::service5::StringOutput& item);
bool Service5StringOutputWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service5::StringOutput& item);

}} // namespace tl2::details

