#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/service1.Value.h"

namespace tl2 { namespace details { 

void BuiltinTuple3Service1ValueReset(std::array<::tl2::service1::Value, 3>& item);

bool BuiltinTuple3Service1ValueWriteJSON(std::ostream & s, const std::array<::tl2::service1::Value, 3>& item);
bool BuiltinTuple3Service1ValueRead(::basictl::tl_istream & s, std::array<::tl2::service1::Value, 3>& item);
bool BuiltinTuple3Service1ValueWrite(::basictl::tl_ostream & s, const std::array<::tl2::service1::Value, 3>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorService1ValueReset(std::vector<::tl2::service1::Value>& item);

bool BuiltinVectorService1ValueWriteJSON(std::ostream & s, const std::vector<::tl2::service1::Value>& item);
bool BuiltinVectorService1ValueRead(::basictl::tl_istream & s, std::vector<::tl2::service1::Value>& item);
bool BuiltinVectorService1ValueWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service1::Value>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void Service1ValueReset(::tl2::service1::Value& item) noexcept;

bool Service1ValueWriteJSON(std::ostream & s, const ::tl2::service1::Value& item) noexcept;
bool Service1ValueReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Value& item) noexcept;
bool Service1ValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Value& item) noexcept;

}} // namespace tl2::details

namespace tl2 { namespace details { 

bool Service1ValueBoxedMaybeWriteJSON(std::ostream & s, const std::optional<::tl2::service1::Value>& item);

bool Service1ValueBoxedMaybeReadBoxed(::basictl::tl_istream & s, std::optional<::tl2::service1::Value>& item);
bool Service1ValueBoxedMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<::tl2::service1::Value>& item);


}} // namespace tl2::details

