#pragma once

#include "../../../basics/basictl.h"
#include "../types/pkg2.foo.h"

namespace tl2 { namespace details { 

void Pkg2FooReset(::tl2::pkg2::Foo& item);

bool Pkg2FooWriteJSON(std::ostream& s, const ::tl2::pkg2::Foo& item);
bool Pkg2FooRead(::basictl::tl_istream & s, ::tl2::pkg2::Foo& item);
bool Pkg2FooWrite(::basictl::tl_ostream & s, const ::tl2::pkg2::Foo& item);
bool Pkg2FooReadBoxed(::basictl::tl_istream & s, ::tl2::pkg2::Foo& item);
bool Pkg2FooWriteBoxed(::basictl::tl_ostream & s, const ::tl2::pkg2::Foo& item);

}} // namespace tl2::details

