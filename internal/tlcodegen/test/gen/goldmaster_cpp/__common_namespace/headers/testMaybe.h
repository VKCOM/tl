// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/testMaybe.h"

namespace tlgen { namespace details { 

void TestMaybeReset(::tlgen::TestMaybe& item) noexcept;

bool TestMaybeWriteJSON(std::ostream& s, const ::tlgen::TestMaybe& item) noexcept;
bool TestMaybeRead(::tlgen::basictl::tl_istream & s, ::tlgen::TestMaybe& item) noexcept; 
bool TestMaybeWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::TestMaybe& item) noexcept;
bool TestMaybeReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::TestMaybe& item);
bool TestMaybeWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::TestMaybe& item);

}} // namespace tlgen::details

