#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "pkg2/types/pkg2.foo.h"

namespace tl2 { namespace details { 

void Pkg2FooReset(::tl2::pkg2::Foo& item) noexcept;

bool Pkg2FooWriteJSON(std::ostream& s, const ::tl2::pkg2::Foo& item) noexcept;
bool Pkg2FooRead(::basictl::tl_istream & s, ::tl2::pkg2::Foo& item) noexcept; 
bool Pkg2FooWrite(::basictl::tl_ostream & s, const ::tl2::pkg2::Foo& item) noexcept;
bool Pkg2FooReadBoxed(::basictl::tl_istream & s, ::tl2::pkg2::Foo& item);
bool Pkg2FooWriteBoxed(::basictl::tl_ostream & s, const ::tl2::pkg2::Foo& item);

}} // namespace tl2::details

