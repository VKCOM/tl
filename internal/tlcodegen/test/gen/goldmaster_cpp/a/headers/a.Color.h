// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "a/types/a.Color.h"

namespace tl2 { namespace details { 

void AColorReset(::tl2::a::Color& item) noexcept;

bool AColorWriteJSON(std::ostream & s, const ::tl2::a::Color& item) noexcept;
bool AColorReadBoxed(::basictl::tl_istream & s, ::tl2::a::Color& item) noexcept;
bool AColorWriteBoxed(::basictl::tl_ostream & s, const ::tl2::a::Color& item) noexcept;

}} // namespace tl2::details

namespace tl2 { namespace details { 

bool AColorBoxedMaybeWriteJSON(std::ostream & s, const std::optional<::tl2::a::Color>& item);

bool AColorBoxedMaybeReadBoxed(::basictl::tl_istream & s, std::optional<::tl2::a::Color>& item);
bool AColorBoxedMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<::tl2::a::Color>& item);


}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorAColorReset(std::vector<::tl2::a::Color>& item);

bool BuiltinVectorAColorWriteJSON(std::ostream & s, const std::vector<::tl2::a::Color>& item);
bool BuiltinVectorAColorRead(::basictl::tl_istream & s, std::vector<::tl2::a::Color>& item);
bool BuiltinVectorAColorWrite(::basictl::tl_ostream & s, const std::vector<::tl2::a::Color>& item);

}} // namespace tl2::details

