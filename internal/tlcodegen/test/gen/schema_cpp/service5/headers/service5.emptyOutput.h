#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service5/types/service5.emptyOutput.h"

namespace tl2 { namespace details { 

void Service5EmptyOutputReset(::tl2::service5::EmptyOutput& item) noexcept;

bool Service5EmptyOutputWriteJSON(std::ostream& s, const ::tl2::service5::EmptyOutput& item) noexcept;
bool Service5EmptyOutputRead(::basictl::tl_istream & s, ::tl2::service5::EmptyOutput& item) noexcept; 
bool Service5EmptyOutputWrite(::basictl::tl_ostream & s, const ::tl2::service5::EmptyOutput& item) noexcept;
bool Service5EmptyOutputReadBoxed(::basictl::tl_istream & s, ::tl2::service5::EmptyOutput& item);
bool Service5EmptyOutputWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service5::EmptyOutput& item);

}} // namespace tl2::details

